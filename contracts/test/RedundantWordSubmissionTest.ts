import { expect } from "chai";
import { ethers } from "hardhat";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
} from "../utils/crypto";
import { getRandomWord } from "../utils/dictionary";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";

describe("Redundant Word Submission", function () {
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

  async function preparePhrase() {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);
    return pid;
  }

  it("should emit RedundantWordRequested after each word submission", async function () {
    const pid = await preparePhrase();
    const selected = await aaa.connect(owner).getSelectedNodes(pid);

    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await expect(aaa.connect(nodeSigner).submitEncryptedWord(pid, word))
      .to.emit(aaa, "RedundantWordRequested")
      .withArgs(pid, 0, nodeSigner.address, anyValue);
  });

  it("should revert if non-UIP node submits redundant word", async function () {
    const pid = await preparePhrase();
    const encWord = ethers.toUtf8Bytes("redundant");

    await expect(
      aaa.connect(nonNode).submitRedundantWord(pid, encWord, 0, nodePk)
    ).to.be.revertedWith("Not UIP node");
  });

  it("should revert if redundant word is empty", async function () {
    const pid = await preparePhrase();
    const selected = await aaa.connect(owner).getSelectedNodes(pid);

    // Force a redundant node from selection
    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];

    // submit the original word to generate redundancy events
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));
    await aaa.connect(nodeSigner).submitEncryptedWord(pid, word);

    const redundantEvents = await aaa.queryFilter(
      aaa.filters.RedundantWordRequested(pid)
    );
    const redundantNode = redundantEvents[0].args.toNode;

    const redundantIndex = nodeAddrs.findIndex(
      (addr) => addr === redundantNode
    );
    const redundantSigner = nodes[redundantIndex];

    await expect(
      aaa
        .connect(redundantSigner)
        .submitRedundantWord(pid, new Uint8Array(), 0, nodePk)
    ).to.be.revertedWith("empty");
  });

  it("should allow a requested node to submit redundant word", async function () {
    const pid = await preparePhrase();
    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await aaa.connect(nodeSigner).submitEncryptedWord(pid, word);

    const redundantEvents = await aaa.queryFilter(
      aaa.filters.RedundantWordRequested(pid)
    );
    const redundantNode = redundantEvents[0].args.toNode;

    const redundantIndex = nodeAddrs.findIndex(
      (addr) => addr === redundantNode
    );
    const redundantSigner = nodes[redundantIndex];
    const redundantWord = await getRandomWord().then((w) =>
      encryptWithKey(w, nodePk)
    );

    await expect(
      aaa
        .connect(redundantSigner)
        .submitRedundantWord(pid, redundantWord, 0, nodePk)
    )
      .to.emit(aaa, "RedundantWordSubmitted")
      .withArgs(
        pid,
        anyValue,
        redundantSigner.address,
        ethers.keccak256(redundantWord)
      );
  });

  it("should revert if a redundant node submits twice", async function () {
    const pid = await preparePhrase();
    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));

    await aaa.connect(nodeSigner).submitEncryptedWord(pid, word);

    const redundantEvents = await aaa.queryFilter(
      aaa.filters.RedundantWordRequested(pid)
    );
    const redundantNode = redundantEvents[0].args.toNode;

    const redundantIndex = nodeAddrs.findIndex(
      (addr) => addr === redundantNode
    );
    const redundantSigner = nodes[redundantIndex];
    const redundantWord = await getRandomWord().then((w) =>
      encryptWithKey(w, nodePk)
    );

    await aaa
      .connect(redundantSigner)
      .submitRedundantWord(pid, redundantWord, 0, nodePk);

    await expect(
      aaa
        .connect(redundantSigner)
        .submitRedundantWord(pid, redundantWord, 0, nodePk)
    ).to.be.revertedWith("already submitted");
  });

  it("should retrieve redundant words via getter", async function () {
    const pid = await preparePhrase();
    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    const firstNodeAddr = selected[0];
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === firstNodeAddr);
    const nodeSigner = nodes[nodeIndex];

    // Submit an original word to trigger redundant requests
    const word = await getRandomWord().then((w) => encryptWithKey(w, pk));
    const tx = await aaa.connect(nodeSigner).submitEncryptedWord(pid, word);
    const receipt = await tx.wait();

    const wordSubmittedEvent = receipt.logs
      .map((log: any) => {
        try {
          return aaa.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .find((log: any) => log?.name === "WordSubmitted");

    expect(wordSubmittedEvent).to.not.be.undefined;
    const index = Number(wordSubmittedEvent!.args.index);

    // Pick the first redundant node
    const redundantEvents = await aaa.queryFilter(
      aaa.filters.RedundantWordRequested(pid)
    );
    const redundantNodeAddr = redundantEvents[0].args.toNode;
    const redundantNodeIndex = nodeAddrs.findIndex(
      (addr) => addr === redundantNodeAddr
    );
    const redundantSigner = nodes[redundantNodeIndex];

    const redundantWord = await getRandomWord().then((w) =>
      encryptWithKey(w, nodePk)
    );

    await aaa
      .connect(redundantSigner)
      .submitRedundantWord(pid, redundantWord, 0, nodePk);

    const [words, nodePKs] = await aaa.getRedundantWords(pid, index);
    expect(words.length).to.be.greaterThan(0);
    expect(nodePKs.length).to.equal(words.length);
    expect(ethers.keccak256(words[0])).to.equal(
      ethers.keccak256(redundantWord)
    );
  });
});
