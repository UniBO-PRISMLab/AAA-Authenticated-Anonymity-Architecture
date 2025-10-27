import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { ethers } from "hardhat";

import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
  decryptWithKey,
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
  let nodeAddrs: string[];
  let pk: Buffer;
  let nodePk: Buffer;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 2);
    nonNode = signers[WORDS_NEEDED + 2];

    const Factory = await ethers.getContractFactory("AAAContract");
    nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

    pk = await generatePublicKey();
    nodePk = await generatePublicKey();

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

  describe("Seed Phrase Protocol", function () {
    it("should start seed phrase generation", async function () {
      // The PID is a 32-byte value
      const pid = generateRandomBytes(32);

      // The public key is a 2048 bit RSA PKCS#8 key
      const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);

      // Verify that the SeedPhraseProtocolInitiated event was emitted with the correct PID
      await expect(tx)
        .to.emit(aaa, "SeedPhraseProtocolInitiated")
        .withArgs(pid);
    });

    it("should not start seed phrase generation if already started", async function () {
      const pid = generateRandomBytes(32);
      const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);
      await expect(tx)
        .to.emit(aaa, "SeedPhraseProtocolInitiated")
        .withArgs(pid);
      await expect(
        aaa.seedPhraseGenerationProtocol(pid, pk)
      ).to.be.revertedWith("already started");
    });

    it("should emit WordRequested events", async function () {
      const pid = generateRandomBytes(32);
      const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);

      // Ignoring the second argument (node address)
      await expect(tx)
        .to.emit(aaa, "WordRequested")
        .withArgs(pid, anyValue, pk);
    });
  });

  describe("Word Submission", function () {
    it("should not allow non-UIP nodes to submit words", async function () {
      const pid = generateRandomBytes(32);
      await aaa.seedPhraseGenerationProtocol(pid, pk);
      const word = await getRandomWord().then(ethers.toUtf8Bytes);
      await expect(
        aaa.connect(nonNode).submitEncryptedWord(pid, word, nodePk)
      ).to.be.revertedWith("Not UIP node");
    });

    it("should not allow a UIP node to submit more than one word", async function () {
      const pid = generateRandomBytes(32);
      await aaa.seedPhraseGenerationProtocol(pid, pk);
      const selected = await aaa.connect(owner).getSelectedNodes(pid);

      const nodeAddr = selected[0];
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      expect(nodeIndex).to.be.greaterThan(-1);
      const nodeSigner = nodes[nodeIndex];

      // Word is encrypted with user's public key, encoded in base64 and then converted to bytes
      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

      // Connect to one of the selected node and submit the word
      await expect(
        await aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
      )
        .to.emit(aaa, "WordSubmitted")
        .withArgs(pid, nodeAddr, ethers.keccak256(word));

      // Connect again to the same node and try to submit another word
      await expect(
        aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
      ).to.be.revertedWith("already submitted");
    });

    it("should not allow a UIP node to submit a empty word", async function () {
      const pid = generateRandomBytes(32);

      await aaa.seedPhraseGenerationProtocol(pid, pk);
      const selected = await aaa.connect(owner).getSelectedNodes(pid);

      const nodeAddr = selected[0];
      const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      expect(nodeIndex).to.be.greaterThan(-1);
      const nodeSigner = nodes[nodeIndex];

      await expect(
        aaa
          .connect(nodeSigner)
          .submitEncryptedWord(pid, new Uint8Array(), nodePk)
      ).to.be.revertedWith("empty");
    });

    it("should not not allow a non selected UIP node to submit a word", async function () {
      const pid = generateRandomBytes(32);
      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

      const selectedLower = selected.map((a: string) => a.toLowerCase());
      const nonSelectedAddr = nodeAddrs.find(
        (addr) => !selectedLower.includes(addr.toLowerCase())
      );

      expect(nonSelectedAddr, "Expected at least one non-selected node").to.not
        .be.undefined;

      const nonSelectedIndex = nodeAddrs.indexOf(nonSelectedAddr);
      const nodeSigner = nodes[nonSelectedIndex];

      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

      await expect(
        aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
      ).to.be.revertedWith("not selected");
    });

    it("should allow selected UIP nodes to submit words and trigger SIDEncryptionRequested", async function () {
      const pid = generateRandomBytes(32);
      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      for (let i = 0; i < selected.length - 1; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];
        const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(word));
      }

      // Last node submission should trigger SIDEncryptionRequested
      const lastNodeAddr = selected[selected.length - 1];
      const lastNodeIndex = nodeAddrs.findIndex(
        (addr) => addr === lastNodeAddr
      );
      expect(lastNodeIndex).to.be.greaterThan(-1);
      const lastNodeSigner = nodes[lastNodeIndex];
      const lastWord = await getRandomWord().then((w) => encryptWithKey(w, pk));

      await expect(
        aaa.connect(lastNodeSigner).submitEncryptedWord(pid, lastWord, pk)
      )
        .to.emit(aaa, "SIDEncryptionRequested")
        .withArgs(pid, anyValue, anyValue, pk);
    });
  });

  describe("SID Encryption", function () {
    it("should not allow non-UIP nodes to submit encrypted SID", async function () {
      const pid = generateRandomBytes(32);

      await expect(
        aaa.connect(nonNode).submitEncryptedSID(pid, new Uint8Array())
      ).to.be.revertedWith("Not UIP node");
    });

    it("should not allow UIP nodes to submit encrypted SID if protocol is not started", async function () {
      const pid = generateRandomBytes(32);

      await expect(
        aaa.connect(nodes[0]).submitEncryptedSID(pid, new Uint8Array())
      ).to.be.revertedWith("not started");
    });

    it("should not allow UIP nodes to submit encrypted SID if was not selected", async function () {
      const pid = generateRandomBytes(32);
      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;

      // Encrypt SID with user's public key
      const pem = Buffer.from(userPK.slice(2), "hex");
      const encSID = await encryptWithKey(sid, pem);

      // Submit from a non-selected node
      const nonSelectedAddr = nodeAddrs.find(
        (addr) => !selected.includes(addr)
      );
      const nonSelectedIndex = nodeAddrs.indexOf(nonSelectedAddr!);
      const nonSelectedSigner = nodes[nonSelectedIndex];
      await expect(
        aaa.connect(nonSelectedSigner).submitEncryptedSID(pid, encSID)
      ).to.be.revertedWith("not selected");
    });

    it("should allow selected UIP node to submit the encrypted SID", async function () {
      const pid = generateRandomBytes(32);

      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];

      // Encrypt SID with user's public key
      const pem = Buffer.from(userPK.slice(2), "hex");
      const encSID = await encryptWithKey(sid, pem);

      await expect(aaa.connect(selectedSigner).submitEncryptedSID(pid, encSID))
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pid, anyValue, anyValue, sid);
    });
  });

  describe("PID Encryption", function () {
    it("should not allow non-UIP nodes to submit encrypted PID", async function () {
      const pid = generateRandomBytes(32);
      await expect(
        aaa
          .connect(nonNode)
          .submitEncryptedPID(pid, new Uint8Array(32), new Uint8Array())
      ).to.be.revertedWith("Not UIP node");
    });

    it("should allow selected UIP node to submit the encrypted PID", async function () {
      const pid = generateRandomBytes(32);

      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];

      // Encrypt SID with user's public key
      const pem = Buffer.from(userPK.slice(2), "hex");
      const encSID = await encryptWithKey(sid, pem);

      await expect(aaa.connect(selectedSigner).submitEncryptedSID(pid, encSID))
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pid, anyValue, anyValue, sid);

      const pidEncFilter = aaa.filters.PIDEncryptionRequested(pid);
      const pidEncEvents = await aaa.queryFilter(pidEncFilter);
      const pidSelectedNode = pidEncEvents[0].args?.node;
      const pidSelectedIndex = nodeAddrs.findIndex(
        (addr) => addr === pidSelectedNode
      );
      const pidSelectedSigner = nodes[pidSelectedIndex];
      const pidReceived = pidEncEvents[0].args?.sid;
      const sidReceived = pidEncEvents[0].args?.sid;

      const symK = pidEncEvents[0].args?.symK;

      // Remove the "0x" prefix if present
      const symEncPID = await encryptWithSymK(
        Buffer.from(symK.slice(2), "hex"),
        pidReceived
      );

      await expect(
        aaa
          .connect(pidSelectedSigner)
          .submitEncryptedPID(pidReceived, sidReceived, symEncPID)
      )
        .to.emit(aaa, "PhraseComplete")
        .withArgs(pidReceived, anyValue);
    });
  });

  describe("SAC", function () {
    it("should allow UIP nodes to submit SAC codes", async function () {
      const sacCodes = [123456, 234567, 345678];

      for (const sac of sacCodes) {
        await expect(aaa.connect(nodes[0]).submitSAC(sac)).to.not.be.reverted;
      }
    });

    it("should not allow non-UIP nodes to submit SAC codes", async function () {
      const sac = 123456;
      await expect(aaa.connect(nonNode).submitSAC(sac)).to.be.revertedWith(
        "Not UIP node"
      );
    });

    it("should store SAC records", async function () {
      const sac = 123456;
      await aaa.connect(nodes[0]).submitSAC(sac);
      const pk = await generatePublicKey();
      await expect(aaa.connect(nodes[0]).submitSACRecord(sac, pk)).to.not.be
        .reverted;
    });

    it("should not allow storing multiple SAC for the same public key", async function () {
      const sac = 123456;
      await aaa.connect(nodes[0]).submitSAC(sac);
      const pk = await generatePublicKey();
      await expect(aaa.connect(nodes[0]).submitSACRecord(sac, pk)).to.not.be
        .reverted;
      await expect(
        aaa.connect(nodes[0]).submitSACRecord(sac, pk)
      ).to.be.revertedWith("already stored");
    });
  });

  describe("Getters", function () {
    it("should retrieve the SID", async function () {
      const pid = generateRandomBytes(32);
      const keypair = await generateKeypair();
      const pk = keypair.public_key;

      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pid);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];
      console.log("Received SID:", sid);
      console.log(
        "Received SID base64:",
        Buffer.from(sid.slice(2)).toString("base64")
      );

      // Encrypt SID with user's public key
      const pem = Buffer.from(userPK.slice(2), "hex");
      const encSID = await encryptWithKey(sid, pem);

      await expect(aaa.connect(selectedSigner).submitEncryptedSID(pid, encSID))
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pid, anyValue, anyValue, sid);

      const pidEncFilter = aaa.filters.PIDEncryptionRequested(pid);
      const pidEncEvents = await aaa.queryFilter(pidEncFilter);
      const pidSelectedNode = pidEncEvents[0].args?.node;
      const pidSelectedIndex = nodeAddrs.findIndex(
        (addr) => addr === pidSelectedNode
      );
      const pidSelectedSigner = nodes[pidSelectedIndex];
      const pidReceived = pidEncEvents[0].args?.sid;
      const sidReceived = pidEncEvents[0].args?.sid;

      const symK = pidEncEvents[0].args?.symK;

      const symEncPID = await encryptWithSymK(
        Buffer.from(symK.slice(2), "hex"),
        pidReceived
      );

      await expect(
        aaa
          .connect(pidSelectedSigner)
          .submitEncryptedPID(pidReceived, sidReceived, symEncPID)
      )
        .to.emit(aaa, "PhraseComplete")
        .withArgs(pidReceived, anyValue);

      const receivedSID = await aaa.getSID(pid);
      const encryptedBuf = Buffer.from(receivedSID.slice(2), "hex");
      const decrypted = await decryptWithKey(encryptedBuf, keypair);

      expect(decrypted).to.equal(sid);
    });

    it("should retrieve the enncrypted words and decrypt them", async function () {
      const pid = generateRandomBytes(32);

      const keypair = await generateKeypair();
      const pk = keypair.public_key;

      await aaa.seedPhraseGenerationProtocol(pid, pk);

      const selected = await aaa.connect(owner).getSelectedNodes(pid);
      let originalWords: string[] = [];
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const word = await getRandomWord();
        originalWords.push(word);
        const wordBuffer = await encryptWithKey(word, pk);

        await expect(
          aaa.connect(nodeSigner).submitEncryptedWord(pid, wordBuffer, nodePk)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pid, nodeSigner.address, ethers.keccak256(wordBuffer));
      }

      const encWordsBytes = aaa.getWords(pid);
      for (let i = 0; i < encWordsBytes.length; i++) {
        const encWordBuf: BytesLike = encWordsBytes[i];
        let buf: Buffer;
        if (typeof encWordBuf === "string") {
          buf = Buffer.from(encWordBuf.slice(2), "hex");
        } else {
          buf = Buffer.from(encWordBuf);
        }
        const decryptedWord = await decryptWithKey(buf, keypair);
        expect(decryptedWord).to.equal(originalWords[i]);
      }
    });
  });
});
