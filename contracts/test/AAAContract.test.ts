import { expect } from "chai";
import { ethers } from "hardhat";
import { getRandomWord } from "./utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
} from "./utils/crypto";

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

  describe("Deployment", function () {
    it("should deploy with correct parameters", async function () {
      expect(await aaa.WORDS_NEEDED()).to.equal(WORDS_NEEDED);
      expect(await aaa.REDUNDANCY_M()).to.equal(REDUNDANCY_M);
      expect(await aaa.owner()).to.equal(await owner.getAddress());
    });

    it("should revert if deployed with too few nodes", async function () {
      const Factory = await ethers.getContractFactory("AAAContract");
      const oneNode = [await nodes[0].getAddress()];
      await expect(Factory.deploy(oneNode, 2, REDUNDANCY_M)).to.be.revertedWith(
        "too few nodes"
      );
    });
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

      const selected = await aaa.getSelectedNodes(pid);
      expect(selected.length).to.equal(WORDS_NEEDED);

      const storedPK = await aaa.getUserPK(pid);
      expect(storedPK).to.equal(ethers.hexlify(pk));
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

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

      await expect(
        aaa.connect(nonNode).submitEncryptedWord(pid, word)
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

      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, new Uint8Array())
      ).to.be.revertedWith("empty");
    });

    it("should emit WordSubmitted and RedundancyRequested", async function () {
      const pid = ethers.keccak256(ethers.toUtf8Bytes("TEST_PID"));
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);

      const selected = await aaa.getSelectedNodes(pid);
      const nodeAddr = await nodes[0].getAddress();
      const nodeIndex = selected.findIndex((addr: string) => addr === nodeAddr);

      await expect(aaa.connect(nodes[0]).submitEncryptedWord(pid, word))
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

      await aaa.connect(nodes[0]).submitEncryptedWord(pid, word);

      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, word)
      ).to.be.revertedWith("already submitted");
    });

    it("should emit PhraseComplete when all nodes have submitted", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      for (let i = 0; i < WORDS_NEEDED - 1; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w);
      }

      const last = await getRandomWord().then(ethers.toUtf8Bytes);
      await expect(
        aaa.connect(nodes[WORDS_NEEDED - 1]).submitEncryptedWord(pid, last)
      ).to.emit(aaa, "PhraseComplete");
    });

    it("should revert submitting word after phrase finalized", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w);
      }

      const extra = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then(ethers.hexlify);
      await expect(
        aaa.connect(nodes[0]).submitEncryptedWord(pid, extra)
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

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w);
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
    it("should return stored words and keys correctly", async function () {
      const pid = ethers.keccak256(
        ethers.toUtf8Bytes(generateRandomBytes(32).toString("base64"))
      );
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));
      await aaa.seedPhraseGenerationProtocol(pid, pkBytes);

      for (let i = 0; i < WORDS_NEEDED; i++) {
        const w = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then(ethers.hexlify);
        await aaa.connect(nodes[i]).submitEncryptedWord(pid, w);
      }

      const selected = await aaa.getSelectedNodes(pid);
      expect(selected.length).to.equal(WORDS_NEEDED);

      const originals = await aaa.getOriginalEncryptedWords(pid);
      expect(originals.length).to.equal(WORDS_NEEDED);

      const sid = await aaa.getSID(pid);
      const symK = await aaa.getSymK(pid);
      expect(sid).to.not.equal(ethers.ZeroHash);
      expect(symK).to.equal(sid);
    });
  });
});
