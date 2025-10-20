import { expect } from "chai";
import { ethers } from "hardhat";
import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
  decryptWithKey,
  recoverPublicKey,
  generateKeypair,
  encryptWithSymK,
} from "../utils/crypto";
import { BytesLike } from "ethers";

describe("AAAContract", function () {
  const WORDS_NEEDED = 4;
  const REDUNDANCY_M = 2;

  let owner: any;
  let nodes: any[];
  let nonNode: any;
  let aaa: any;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 1);
    nonNode = signers[WORDS_NEEDED + 2];

    const Factory = await ethers.getContractFactory("AAAContract");
    const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

  describe("Seed Phrase Protocol Requested", function () {
    it("should initiate seed phrase and emit the event", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey()
        .then((b) => b.toString("base64"))
        .then(ethers.toUtf8Bytes);

      const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);

      await expect(tx)
        .to.emit(aaa, "SeedPhraseProtocolInitiated")
        .withArgs(pid);
    });

    it("should revert if already started for same pid", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey()
        .then((b) => b.toString("base64"))
        .then(ethers.toUtf8Bytes);
      await aaa.seedPhraseGenerationProtocol(pid, pk);
      await expect(
        aaa.seedPhraseGenerationProtocol(pid, pk)
      ).to.be.revertedWith("already started");
    });
  });

  describe("UIP nodes words submission", function () {
    it("should allow only UIP nodes", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

      await expect(
        aaa.connect(nonNode).submitEncryptedWord(pid, word, nodePkBytes)
      ).to.be.revertedWith("Not UIP node");
    });

    it("should reject empty encrypted words", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey()
        .then((b) => b.toString("base64"))
        .then(ethers.toUtf8Bytes);
      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      await expect(
        aaa
          .connect(nodes[0])
          .submitEncryptedWord(pid, new Uint8Array(), nodePkBytes)
      ).to.be.revertedWith("empty");
    });

    it("should emit WordSubmitted and RedundancyRequested", async function () {
      const pid = ethers.keccak256(ethers.toUtf8Bytes("TEST_PID"));
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);

      const selected = await aaa.getSelectedNodes(pid);
      const nodeAddr = await nodes[0].getAddress();
      const nodeIndex = selected.findIndex((addr: string) => addr === nodeAddr);

      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, word, nodePkBytes)
      )
        .to.emit(aaa, "WordSubmitted")
        .withArgs(
          pid,
          await nodes[0].getAddress(),
          nodeIndex,
          ethers.keccak256(word)
        )
        .and.to.emit(aaa, "RedundancyRequested");
    });

    it("should not allow same node to submit twice for same PID", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      await aaa.connect(nodes[0]).submitEncryptedWord(pid, word, nodePkBytes);

      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, word, nodePkBytes)
      ).to.be.revertedWith("already submitted");
    });

    it("should emit SIDEncryptionRequested when all nodes have submitted", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      for (let i = 0; i < WORDS_NEEDED - 1; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const last = await getRandomWord().then(ethers.toUtf8Bytes);
      await expect(
        aaa
          .connect(nodes[WORDS_NEEDED - 1])
          .submitEncryptedWord(pid, last, nodePkBytes)
      ).to.emit(aaa, "SIDEncryptionRequested");
    });

    it("should emit PIDEncryptionRequested after a node submits the encrypted SID", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );

      const keypair = await generateKeypair();
      const publicKey = keypair.public_key;

      const pkBytes = ethers.toUtf8Bytes(publicKey.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, publicKey))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;

      const { buffer } = recoverPublicKey(userPK);

      const encSID = await encryptWithKey(sid, buffer);
      const encSIDB64 = encSID.toString("base64");
      const encSIDBytes = ethers.toUtf8Bytes(encSIDB64);

      await expect(
        aaa
          .connect(await ethers.getSigner(selectedNode))
          .storeEncryptedSID(pid, encSIDBytes, {
            gasLimit: 5_000_000,
          })
      ).to.emit(aaa, "PIDEncryptionRequested");
    });

    it("should finalize phrase after all words, SID and PID are encrypted and submitted", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );

      const keypair = await generateKeypair();
      const publicKey = keypair.public_key;

      const pkBytes = ethers.toUtf8Bytes(publicKey.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, publicKey))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;

      const { buffer } = recoverPublicKey(userPK);

      const encSID = await encryptWithKey(sid, buffer);
      const encSIDB64 = encSID.toString("base64");
      const encSIDBytes = ethers.toUtf8Bytes(encSIDB64);

      await expect(
        aaa
          .connect(await ethers.getSigner(selectedNode))
          .storeEncryptedSID(pid, encSIDBytes, {
            gasLimit: 5_000_000,
          })
      ).to.emit(aaa, "PIDEncryptionRequested");

      const encPID = await encryptWithSymK(pid);
      const encPIDB64 = encPID.toString("base64");
      const encPIDBytes = ethers.toUtf8Bytes(encPIDB64);

      await expect(
        aaa
          .connect(await ethers.getSigner(selectedNode))
          .storeEncryptedPID(pid, encPIDBytes, {
            gasLimit: 5_000_000,
          })
      ).to.emit(aaa, "PhraseComplete");
    });

    it("should revert submitting word after phrase finalized", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const extra = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);
      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, extra, nodePkBytes)
      ).to.be.revertedWith("done");
    });
  });

  describe("Redundancy", function () {
    it("should allow UIP node to submit redundancy once per index", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const redundancy = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);
      await expect(
        aaa.connect(nodes[0]).submitRedundantEncryptedWord(pid, 1, redundancy)
      ).to.emit(aaa, "RedundantWordSubmitted");

      await expect(
        aaa.connect(nodes[0]).submitRedundantEncryptedWord(pid, 1, redundancy)
      ).to.be.revertedWith("already submitted");
    });
  });

  describe("Getters", function () {
    it("should return the list of words for a PID and decrypt them", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const keypair = await generateKeypair();
      const publicKey = keypair.public_key;
      const pkBytes = ethers.toUtf8Bytes(publicKey.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      const originalWords: string[] = [];
      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord();
        originalWords.push(w);
        const encW = await encryptWithKey(w, publicKey).then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, encW, nodePkBytes);
      }

      const storedWords = await aaa.getWords(pid);
      const storedWordsHex = storedWords.map((w: BytesLike) =>
        ethers.hexlify(w)
      );

      expect(storedWordsHex.length).to.equal(WORDS_NEEDED);

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const dec = await decryptWithKey(
          Buffer.from(storedWordsHex[i].slice(2), "hex"),
          keypair
        );
        expect(dec).to.equal(originalWords[i]);
      }
    });

    it("should return the encrypted SID after phrase completion", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );

      const keypair = await generateKeypair();
      const publicKey = keypair.public_key;

      const pkBytes = ethers.toUtf8Bytes(publicKey.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, publicKey))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w, nodePkBytes);
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;

      const { buffer } = recoverPublicKey(userPK);

      const encSID = await encryptWithKey(sid, buffer);
      const encSIDB64 = encSID.toString("base64");
      const encSIDBytes = ethers.toUtf8Bytes(encSIDB64);

      await expect(
        aaa
          .connect(await ethers.getSigner(selectedNode))
          .storeEncryptedSID(pid, encSIDBytes, {
            gasLimit: 5_000_000,
          })
      ).to.emit(aaa, "PIDEncryptionRequested");

      const storedEncSIDHex = await aaa.getSID(pid);
      const storedEncSIDBase64 = Buffer.from(
        storedEncSIDHex.slice(2),
        "hex"
      ).toString("utf8");
      const storedEncSIDBuffer = Buffer.from(storedEncSIDBase64, "base64");

      expect(storedEncSIDHex).to.equal(ethers.hexlify(encSIDBytes));

      const dec1 = await decryptWithKey(encSID, keypair);
      const dec2 = await decryptWithKey(storedEncSIDBuffer, keypair);
      expect(dec2).to.equal(dec1);
    });
  });
});
