import { expect } from "chai";
import { ethers } from "hardhat";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";

import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
  decryptWithKey,
  recoverPublicKey,
  generateKeypair,
  encryptSym,
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

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 2);
    nonNode = signers[WORDS_NEEDED + 2];

    const Factory = await ethers.getContractFactory("AAAContract");
    nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

  describe("Seed Phrase Protocol", function () {
    it("should start seed phrase generation", async function () {
      // The PID is a 32-byte value encoded in base64
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");

      // The public key is a 2048 bit RSA PKCS#8 key encoded in base64
      const pk = await generatePublicKey().then((b) => b.toString("base64"));

      const pkBytes = ethers.toUtf8Bytes(pk);

      const tx = await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      // Verify that the SeedPhraseProtocolInitiated event was emitted with the correct PID
      await expect(tx)
        .to.emit(aaa, "SeedPhraseProtocolInitiated")
        .withArgs(pidBytes);

      // Convert the pidBytes back to base64 for comparison
      const pidBase64 = Buffer.from(pidBytes).toString("base64");
      expect(pidBase64).to.equal(pid);
    });

    it("should not start seed phrase generation if already started", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey().then((b) => b.toString("base64"));
      const pkBytes = ethers.toUtf8Bytes(pk);

      const tx = await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      await expect(tx)
        .to.emit(aaa, "SeedPhraseProtocolInitiated")
        .withArgs(pidBytes);

      await expect(
        aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes)
      ).to.be.revertedWith("already started");
    });

    it("should emit WordRequested events", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey().then((b) => b.toString("base64"));
      const pkBytes = ethers.toUtf8Bytes(pk);

      const tx = await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      // Ignoring the second argument (node address)
      await expect(tx)
        .to.emit(aaa, "WordRequested")
        .withArgs(pidBytes, anyValue, pkBytes);
    });
  });

  describe("Word Submission", function () {
    it("should not allow non-UIP nodes to submit words", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey().then((b) => b.toString("base64"));
      const pkBytes = ethers.toUtf8Bytes(pk);

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      const word = await getRandomWord().then(ethers.toUtf8Bytes);

      await expect(
        aaa.connect(nonNode).submitEncryptedWord(pidBytes, word, nodePkBytes)
      ).to.be.revertedWith("Not UIP node");
    });

    it("should not allow a UIP node to submit more than one word", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);

      const nodeAddr = selected[0];
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      expect(nodeIndex).to.be.greaterThan(-1);
      const nodeSigner = nodes[nodeIndex];

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      // Word is encrypted with user's public key, encoded in base64 and then converted to bytes
      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then((enc) => ethers.getBytes(enc));

      // Connect to one of the selected node and submit the word
      await expect(
        await aaa
          .connect(nodeSigner)
          .submitEncryptedWord(pidBytes, word, nodePkBytes)
      )
        .to.emit(aaa, "WordSubmitted")
        .withArgs(pidBytes, nodeAddr, ethers.keccak256(word));

      // Connect again to the same node and try to submit another word
      await expect(
        aaa.connect(nodeSigner).submitEncryptedWord(pidBytes, word, nodePkBytes)
      ).to.be.revertedWith("already submitted");
    });

    it("should not allow a UIP node to submit a empty word", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);

      const nodeAddr = selected[0];
      const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      expect(nodeIndex).to.be.greaterThan(-1);
      const nodeSigner = nodes[nodeIndex];

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      await expect(
        aaa
          .connect(nodeSigner)
          .submitEncryptedWord(pidBytes, new Uint8Array(), nodePkBytes)
      ).to.be.revertedWith("empty");
    });

    it("should not not allow a non selected UIP node to submit a word", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

      const selectedLower = selected.map((a: string) => a.toLowerCase());
      const nonSelectedAddr = nodeAddrs.find(
        (addr) => !selectedLower.includes(addr.toLowerCase())
      );

      expect(nonSelectedAddr, "Expected at least one non-selected node").to.not
        .be.undefined;

      const nonSelectedIndex = nodeAddrs.indexOf(nonSelectedAddr);
      const nodeSigner = nodes[nonSelectedIndex];
      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      const word = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

      await expect(
        aaa.connect(nodeSigner).submitEncryptedWord(pidBytes, word, nodePkBytes)
      ).to.be.revertedWith("not selected");
    });

    it("should allow selected UIP nodes to submit words and trigger SIDEncryptionRequested", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      for (let i = 0; i < selected.length - 1; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, word, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(word));
      }

      // Last node submission should trigger SIDEncryptionRequested
      const lastNodeAddr = selected[selected.length - 1];
      const lastNodeIndex = nodeAddrs.findIndex(
        (addr) => addr === lastNodeAddr
      );
      expect(lastNodeIndex).to.be.greaterThan(-1);
      const lastNodeSigner = nodes[lastNodeIndex];
      const lastNodePk = await generatePublicKey();
      const lastNodePkBytes = ethers.toUtf8Bytes(lastNodePk.toString("base64"));
      const lastWord = await getRandomWord()
        .then((w) => encryptWithKey(w, pk))
        .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

      await expect(
        aaa
          .connect(lastNodeSigner)
          .submitEncryptedWord(pidBytes, lastWord, lastNodePkBytes)
      )
        .to.emit(aaa, "SIDEncryptionRequested")
        .withArgs(pidBytes, anyValue, anyValue, pkBytes);
    });
  });

  describe("SID Encryption", function () {
    it("should not allow non-UIP nodes to submit encrypted SID", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");

      await expect(
        aaa.connect(nonNode).submitEncryptedSID(pidBytes, new Uint8Array())
      ).to.be.revertedWith("Not UIP node");
    });

    it("should not allow UIP nodes to submit encrypted SID if protocol is not started", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");

      await expect(
        aaa.connect(nodes[0]).submitEncryptedSID(pidBytes, new Uint8Array())
      ).to.be.revertedWith("not started");
    });

    it("should not allow UIP nodes to submit encrypted SID if was not selected", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, word, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pidBytes);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;

      // Encrypt SID with user's public key
      const { buffer } = recoverPublicKey(userPK);

      const encSID = await encryptWithKey(sid, buffer);
      const encSIDBytes = ethers.toUtf8Bytes(encSID.toString("base64"));

      // Submit from a non-selected node
      const nonSelectedAddr = nodeAddrs.find(
        (addr) => !selected.includes(addr)
      );
      const nonSelectedIndex = nodeAddrs.indexOf(nonSelectedAddr!);
      const nonSelectedSigner = nodes[nonSelectedIndex];
      await expect(
        aaa.connect(nonSelectedSigner).submitEncryptedSID(pidBytes, encSIDBytes)
      ).to.be.revertedWith("not selected");
    });

    it("should allow selected UIP node to submit the encrypted SID", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, word, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pidBytes);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];

      // Encrypt SID with user's public key
      const { buffer } = recoverPublicKey(userPK);
      const encSID = await encryptWithKey(sid, buffer);
      const encSIDBytes = ethers.toUtf8Bytes(encSID.toString("base64"));

      await expect(
        aaa.connect(selectedSigner).submitEncryptedSID(pidBytes, encSIDBytes)
      )
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pidBytes, anyValue, anyValue, sid);
    });
  });

  describe("PID Encryption", function () {
    it("should not allow non-UIP nodes to submit encrypted PID", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");

      await expect(
        aaa
          .connect(nonNode)
          .submitEncryptedPID(pidBytes, new Uint8Array(32), new Uint8Array())
      ).to.be.revertedWith("Not UIP node");
    });

    it("should allow selected UIP node to submit the encrypted PID", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const pk = await generatePublicKey();
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, word, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pidBytes);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];

      // Encrypt SID with user's public key
      const { buffer } = recoverPublicKey(userPK);
      const encSID = await encryptWithKey(sid, buffer);

      await expect(
        aaa.connect(selectedSigner).submitEncryptedSID(pidBytes, encSID)
      )
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pidBytes, anyValue, anyValue, sid);

      const pidEncFilter = aaa.filters.PIDEncryptionRequested(pidBytes);
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
        pid
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

  describe("Getters", function () {
    it("should retrieve the SID", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");
      const keypair = await generateKeypair();
      const pk = keypair.public_key;
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord()
          .then((w) => encryptWithKey(w, pk))
          .then((enc) => ethers.toUtf8Bytes(enc.toString("base64")));

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, word, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(word));
      }

      const filter = aaa.filters.SIDEncryptionRequested(pidBytes);
      const events = await aaa.queryFilter(filter);
      const selectedNode = events[0].args?.node;
      const sid = events[0].args?.sid;
      const userPK = events[0].args?.userPK;
      const selectedIndex = nodeAddrs.findIndex(
        (addr) => addr === selectedNode
      );
      const selectedSigner = nodes[selectedIndex];

      // Encrypt SID with user's public key
      const { buffer } = recoverPublicKey(userPK);
      const encSID = await encryptWithKey(sid, buffer);

      await expect(
        aaa.connect(selectedSigner).submitEncryptedSID(pidBytes, encSID)
      )
        .to.emit(aaa, "PIDEncryptionRequested")
        .withArgs(pidBytes, anyValue, anyValue, sid);

      const pidEncFilter = aaa.filters.PIDEncryptionRequested(pidBytes);
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
        pid
      );

      await expect(
        aaa
          .connect(pidSelectedSigner)
          .submitEncryptedPID(pidReceived, sidReceived, symEncPID)
      )
        .to.emit(aaa, "PhraseComplete")
        .withArgs(pidReceived, anyValue);

      const receivedSID = await aaa.getSID(pidBytes);
      const encryptedBuf = Buffer.from(receivedSID.slice(2), "hex");
      const decrypted = await decryptWithKey(encryptedBuf, keypair);

      expect(decrypted).to.equal(sid);
    });

    it("should retrieve the enncrypted words and decrypt them", async function () {
      const pid = generateRandomBytes(32).toString("base64");
      const pidBytes = Buffer.from(pid, "base64");

      const keypair = await generateKeypair();
      const pk = keypair.public_key;
      const pkBytes = ethers.toUtf8Bytes(pk.toString("base64"));

      await aaa.seedPhraseGenerationProtocol(pidBytes, pkBytes);

      const selected = await aaa.connect(owner).getSelectedNodes(pidBytes);
      let originalWords: string[] = [];
      for (let i = 0; i < selected.length; i++) {
        const nodeAddr = selected[i];
        const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
        expect(nodeIndex).to.be.greaterThan(-1);
        const nodeSigner = nodes[nodeIndex];

        const nodePk = await generatePublicKey();
        const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

        const word = await getRandomWord();
        originalWords.push(word);
        const wordBuffer = await encryptWithKey(word, pk);

        await expect(
          aaa
            .connect(nodeSigner)
            .submitEncryptedWord(pidBytes, wordBuffer, nodePkBytes)
        )
          .to.emit(aaa, "WordSubmitted")
          .withArgs(pidBytes, nodeSigner.address, ethers.keccak256(wordBuffer));
      }

      const encWordsBytes = aaa.getWords(pidBytes);
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
