import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { ethers } from "hardhat";
import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
} from "../utils/crypto";

describe("SID Encryption Submission", function () {
  const WORDS_NEEDED = 4;
  const REDUNDANCY_M = 2;

  let owner: any;
  let nodes: any[];
  let nonNode: any;
  let aaa: any;
  let nodeAddrs: string[];
  let pk: Buffer;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 2);
    nonNode = signers[WORDS_NEEDED + 2];

    const Factory = await ethers.getContractFactory("AAA");
    nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));
    pk = await generatePublicKey();

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

  async function preparePhrase() {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.getSelectedNodes(pid);
    for (const addr of selected) {
      const nodeSigner = nodes[nodeAddrs.indexOf(addr)];
      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));
      await aaa.connect(nodeSigner).submitEncryptedWord(pid, word, pk);
    }
    return pid;
  }

  it("should revert if phrase not started", async function () {
    const pid = generateRandomBytes(32);
    const encSID = ethers.randomBytes(32);
    await expect(
      aaa.connect(nodes[0]).submitEncryptedSID(pid, encSID)
    ).to.be.revertedWith("not started");
  });

  it("should revert if encSID is empty", async function () {
    const pid = await preparePhrase();
    const filter = aaa.filters.SIDEncryptionRequested(pid);
    const events = await aaa.queryFilter(filter);
    expect(events.length).to.be.greaterThan(0);
    const encryptionResp = events[0].args.node;
    const nodeIndex = nodeAddrs.findIndex((addr) => addr === encryptionResp);
    expect(nodeIndex).to.be.greaterThan(-1);
    const nodeSigner = nodes[nodeIndex];
    await expect(
      aaa.connect(nodeSigner).submitEncryptedSID(pid, new Uint8Array())
    ).to.be.revertedWith("empty");
  });

  it("should revert if called by non-selected node", async function () {
    const pid = await preparePhrase();
    const filter = aaa.filters.SIDEncryptionRequested(pid);
    const events = await aaa.queryFilter(filter);
    const encryptionResp = events[0].args.node;
    const wrongNode = nodes.find((n) => n.address !== encryptionResp);
    const encSID = ethers.randomBytes(32);
    await expect(
      aaa.connect(wrongNode).submitEncryptedSID(pid, encSID)
    ).to.be.revertedWith("not selected");
  });

  it("should revert if already stored", async function () {
    const pid = await preparePhrase();
    const filter = aaa.filters.SIDEncryptionRequested(pid);
    const events = await aaa.queryFilter(filter);
    const encryptionResp = events[0].args.node;
    const respSigner = nodes[nodeAddrs.indexOf(encryptionResp)];
    const encSID = ethers.randomBytes(32);
    await aaa.connect(respSigner).submitEncryptedSID(pid, encSID);
    await expect(
      aaa.connect(respSigner).submitEncryptedSID(pid, encSID)
    ).to.be.revertedWith("already stored");
  });

  it("should allow the correct node to submit SID and emit PIDEncryptionRequested", async function () {
    const pid = await preparePhrase();
    const filter = aaa.filters.SIDEncryptionRequested(pid);
    const events = await aaa.queryFilter(filter);
    const encryptionResp = events[0].args.node;
    const respSigner = nodes[nodeAddrs.indexOf(encryptionResp)];
    const encSID = ethers.randomBytes(32);
    await expect(aaa.connect(respSigner).submitEncryptedSID(pid, encSID))
      .to.emit(aaa, "PIDEncryptionRequested")
      .withArgs(pid, encryptionResp, anyValue, anyValue);
    const stored = await aaa.getSID(pid);
    expect(stored).to.equal(ethers.hexlify(encSID));
  });
});
