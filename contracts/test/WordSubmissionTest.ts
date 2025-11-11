import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { ethers } from "hardhat";
import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
} from "../utils/crypto";

describe("Word Submission", function () {
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

    const Factory = await ethers.getContractFactory("AAA");
    nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

    pk = await generatePublicKey();
    nodePk = await generatePublicKey();

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

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
      .withArgs(pid, nodeAddr, ethers.keccak256(word), anyValue);

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
      aaa.connect(nodeSigner).submitEncryptedWord(pid, new Uint8Array(), nodePk)
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

    expect(nonSelectedAddr, "Expected at least one non-selected node").to.not.be
      .undefined;

    const nonSelectedIndex = nodeAddrs.indexOf(nonSelectedAddr);
    const nodeSigner = nodes[nonSelectedIndex];

    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await expect(
      aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
    ).to.be.revertedWith("not selected");
  });

  it("should store submitted word and mark node as submitted", async function () {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.getSelectedNodes(pid);
    const nodeAddr = selected[0];
    const nodeSigner = nodes[nodeAddrs.indexOf(nodeAddr)];

    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk);

    const words = await aaa.getWords(pid);
    expect(words.length).to.equal(1);
    expect(words[0]).to.equal(ethers.hexlify(word));

    // Trying again should revert
    await expect(
      aaa.connect(nodeSigner).submitEncryptedWord(pid, word, nodePk)
    ).to.be.revertedWith("already submitted");
  });

  it("should emit SIDEncryptionRequested only on the final submission", async function () {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.connect(owner).getSelectedNodes(pid);

    for (let i = 0; i < selected.length - 1; i++) {
      const nodeAddr = selected[i];
      const nodeSigner = nodes[nodeAddrs.indexOf(nodeAddr)];
      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

      const tx = await aaa
        .connect(nodeSigner)
        .submitEncryptedWord(pid, word, nodePk);
      const receipt = await tx.wait();

      const sidEvents = receipt.logs
        .map((log: any) => {
          try {
            return aaa.interface.parseLog(log);
          } catch {
            return null;
          }
        })
        .filter((log: any) => log?.name === "SIDEncryptionRequested");

      expect(sidEvents.length).to.equal(0);
    }

    // final submission
    const lastNodeAddr = selected[selected.length - 1];
    const lastNodeSigner = nodes[nodeAddrs.indexOf(lastNodeAddr)];
    const lastWord = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await expect(
      aaa.connect(lastNodeSigner).submitEncryptedWord(pid, lastWord, nodePk)
    ).to.emit(aaa, "SIDEncryptionRequested");
  });

  it("should emit RedundantWordRequested events for each submitted word", async function () {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    const tx = await aaa
      .connect(nodeSigner)
      .submitEncryptedWord(pid, word, nodePk);
    const receipt = await tx.wait();

    const parsed = receipt.logs
      .map((log: any) => {
        try {
          return aaa.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .filter((log: any) => log?.name === "RedundantWordRequested");

    expect(parsed.length).to.equal(REDUNDANCY_M);
  });

  it("should revert if phrase not started", async function () {
    const pid = generateRandomBytes(32);
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await expect(
      aaa.connect(nodes[0]).submitEncryptedWord(pid, word, nodePk)
    ).to.be.revertedWith("not started");
  });

  it("should emit SIDEncryptionRequested with correct user PK", async function () {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.getSelectedNodes(pid);
    for (let i = 0; i < selected.length; i++) {
      const nodeAddr = selected[i];
      const nodeSigner = nodes[nodeAddrs.indexOf(nodeAddr)];
      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

      const tx = await aaa
        .connect(nodeSigner)
        .submitEncryptedWord(pid, word, nodePk);
      const receipt = await tx.wait();

      const sidEvents = receipt.logs
        .map((log: any) => {
          try {
            return aaa.interface.parseLog(log);
          } catch {
            return null;
          }
        })
        .filter((log: any) => log?.name === "SIDEncryptionRequested");

      if (i < selected.length - 1) {
        expect(sidEvents.length).to.equal(0);
      } else {
        expect(sidEvents.length).to.equal(1);
        const args = sidEvents[0].args;
        expect(args.userPK).to.equal(ethers.hexlify(pk));
      }
    }
  });
});
