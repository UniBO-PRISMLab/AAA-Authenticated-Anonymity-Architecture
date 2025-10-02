import { expect } from "chai";
import { ethers } from "hardhat";
import { Signer } from "ethers";
import { AAAContract } from "../types";

describe("AAAContract", function () {
  let aaa: AAAContract;
  let owner: Signer;
  let nodes: Signer[];

  const WORDS_NEEDED = 6;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
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
    const pubKey = ethers.toUtf8Bytes("fakePublicKey");

    await expect(aaa.seedPhraseGenerationProtocol(pid, pubKey))
      .to.emit(aaa, "WordRequestedToUIPNode")
      .withArgs(pid, await nodes[0].getAddress(), pubKey);
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
