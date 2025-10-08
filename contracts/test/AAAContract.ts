import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { AAAContract } from "../types";
import { encrypt, decrypt, loadPublicKey } from "./crypto";
import { getRandomWord } from "./dictionary";

describe("AAAContract", function () {
  let aaa: AAAContract;
  let owner: Signer;
  let nodes: Signer[];
  let trustedNode1: Signer;
  let trustedNode2: Signer;
  let untrustedNode: Signer;

  const WORDS_NEEDED = 6;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];

    trustedNode1 = signers[1];
    trustedNode2 = signers[2];

    untrustedNode = signers[WORDS_NEEDED + 2];

    nodes = signers.slice(1, WORDS_NEEDED + 1);

    const AAAContractFactory = await ethers.getContractFactory("AAAContract");
    const nodeAddresses = await Promise.all(nodes.map((n) => n.getAddress()));
    aaa = (await AAAContractFactory.deploy(
      nodeAddresses
    )) as unknown as AAAContract;
    await aaa.waitForDeployment();
  });

  it("should deploy with enough nodes", async function () {
    expect(await aaa.getAddress()).to.properAddress;
  });

  it("should fail deployment if not enough nodes", async function () {
    const fewerNodes = [await nodes[0].getAddress()];
    const Factory = await ethers.getContractFactory("AAAContract");
    await expect(Factory.deploy(fewerNodes)).to.be.revertedWith(
      "Not enough nodes in the pool"
    );
  });

  it("should emit WordRequestedToUIPNode events", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("testpid"));
    const pubKey = await loadPublicKey().then((buf) =>
      ethers.toUtf8Bytes(buf.toString("base64"))
    );

    await expect(aaa.seedPhraseGenerationProtocol(pid, pubKey))
      .to.emit(aaa, "WordRequestedToUIPNode")
      .withArgs(pid, await nodes[0].getAddress(), pubKey);
  });

  it("should store encrypted words", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("anotherpid"));
    const word = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );

    await aaa.connect(trustedNode1).submitEncryptedWord(pid, word);

    const stored = await aaa.encryptedWords(pid, 0);
    expect(stored).to.equal(ethers.hexlify(word));
  });

  it("should emit PhraseComplete after needed words", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("apid"));
    const finalWord = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );

    for (let i = 0; i < WORDS_NEEDED - 1; i++) {
      const word = await getRandomWord().then((w) =>
        encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
      );
      await aaa.connect(nodes[i]).submitEncryptedWord(pid, word);
    }

    await expect(
      aaa.connect(nodes[WORDS_NEEDED - 1]).submitEncryptedWord(pid, finalWord)
    ).to.emit(aaa, "PhraseComplete");
  });

  it("should not allow untrusted nodes to submit words", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("apid"));
    const word = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );

    await expect(
      aaa.connect(untrustedNode).submitEncryptedWord(pid, word)
    ).to.be.revertedWith("Not an authorized UIP node");
  });

  it("should allow each UIP node to submit exactly one word per PID", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("apid"));
    const anotherPid = ethers.keccak256(ethers.toUtf8Bytes("anotherpid"));
    const word1 = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );
    const word2 = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );
    const word3 = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );

    // First submission from a trusted node should succeed
    await expect(aaa.connect(trustedNode1).submitEncryptedWord(pid, word1)).to
      .not.be.reverted;

    // Second submission from same node for same PID should fail
    await expect(
      aaa.connect(trustedNode1).submitEncryptedWord(pid, word3)
    ).to.be.revertedWith("Node already submitted a word");

    // Submission from different node for same PID should succeed
    await expect(aaa.connect(trustedNode2).submitEncryptedWord(pid, word2)).to
      .not.be.reverted;

    // Submission from same node for different PID should succeed
    await expect(
      aaa.connect(trustedNode1).submitEncryptedWord(anotherPid, word1)
    ).to.not.be.reverted;
  });

  it("should not allow submitting words after phrase is complete", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("apid"));
    const excessWord = await getRandomWord().then((w) =>
      encrypt(w).then((enc) => ethers.toUtf8Bytes(enc.toString("base64")))
    );
    for (let i = 0; i < WORDS_NEEDED; i++) {
      const word = await getRandomWord();
      const encBuffer = await encrypt(word);
      const encBase64 = encBuffer.toString("base64");
      const encBytes = ethers.toUtf8Bytes(encBase64);
      await aaa.connect(nodes[i]).submitEncryptedWord(pid, encBytes);
    }

    await expect(
      aaa.connect(trustedNode1).submitEncryptedWord(pid, excessWord)
    ).to.be.revertedWith("Phrase already completed");
  });

  it("should retrieve encrypted words for a PID", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("apid"));
    const words = [];
    for (let i = 0; i < WORDS_NEEDED; i++) {
      const word = await getRandomWord();
      const encBuffer = await encrypt(word);

      const encBase64 = encBuffer.toString("base64");
      const encBytes = ethers.toUtf8Bytes(encBase64);

      words.push(word);

      await aaa.connect(nodes[i]).submitEncryptedWord(pid, encBytes);
    }

    const retrieved = await aaa.getEncryptedWords(pid);
    for (let i = 0; i < WORDS_NEEDED; i++) {
      const base64String = ethers.toUtf8String(retrieved[i]);
      const ciphertext = Buffer.from(base64String, "base64");
      const decrypted = await decrypt(ciphertext);
      expect(decrypted).to.equal(words[i]);
    }
  });
});
