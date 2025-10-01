const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("AAAContract", function () {
  let AAAContract, aaa, owner, nodes;

  const WORDS_NEEDED = 6;

  beforeEach(async function () {
    [owner, ...addrs] = await ethers.getSigners();

    nodes = addrs.slice(0, WORDS_NEEDED);

    AAAContract = await ethers.getContractFactory("AAAContract");
    aaa = await AAAContract.deploy(nodes.map((n) => n.address));
    await aaa.waitForDeployment();
  });

  it("should deploy with enough nodes", async function () {
    expect(await aaa.getAddress()).to.properAddress;
  });

  it("should fail deployment if not enough nodes", async function () {
    const fewerNodes = [nodes[0].address];
    const Factory = await ethers.getContractFactory("AAAContract");
    await expect(Factory.deploy(fewerNodes)).to.be.revertedWith(
      "Not enough nodes in the pool"
    );
  });

  it("should emit WordRequestedToUIPNode events", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("testpid"));
    const pubKey = ethers.toUtf8Bytes("fakePublicKey");

    await expect(aaa.seedPhraseGenerationProtocol(pid, pubKey))
      .to.emit(aaa, "WordRequestedToUIPNode")
      .withArgs(pid, nodes[0].address, pubKey);
  });

  it("should store encrypted words", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("anotherpid"));
    const word = ethers.toUtf8Bytes("d0gAndC4t");

    await aaa.submitEncryptedWord(pid, word);

    const stored = await aaa.encryptedWords(pid, 0);
    expect(stored).to.equal(ethers.hexlify(word));
  });

  it("should emit PhraseComplete after 6 words", async function () {
    const pid = ethers.keccak256(ethers.toUtf8Bytes("phrasepid"));

    for (let i = 0; i < WORDS_NEEDED - 1; i++) {
      const word = ethers.toUtf8Bytes(`word${i}`);
      await aaa.submitEncryptedWord(pid, word);
    }

    await expect(
      aaa.submitEncryptedWord(pid, ethers.toUtf8Bytes("finalword"))
    ).to.emit(aaa, "PhraseComplete");
  });
});
