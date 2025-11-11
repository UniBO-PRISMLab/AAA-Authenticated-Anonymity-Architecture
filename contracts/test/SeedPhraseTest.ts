import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { ethers } from "hardhat";
import { generatePublicKey, generateRandomBytes } from "../utils/crypto";

describe("Seed Phrase Generation Protocol", function () {
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

  it("should start seed phrase generation", async function () {
    // The PID is a 32-byte value
    const pid = generateRandomBytes(32);

    // The public key is a 2048 bit RSA PKCS#8 key
    const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);

    // Verify that the SeedPhraseProtocolInitiated event was emitted with the correct PID
    await expect(tx).to.emit(aaa, "SeedPhraseProtocolInitiated").withArgs(pid);
  });

  it("should not start seed phrase generation if already started", async function () {
    const pid = generateRandomBytes(32);
    const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);
    await expect(tx).to.emit(aaa, "SeedPhraseProtocolInitiated").withArgs(pid);
    await expect(aaa.seedPhraseGenerationProtocol(pid, pk)).to.be.revertedWith(
      "already started"
    );
  });

  it("should emit WordRequested events", async function () {
    const pid = generateRandomBytes(32);
    const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);

    // Ignoring the second argument (node address)
    await expect(tx).to.emit(aaa, "WordRequested").withArgs(pid, anyValue, pk);
  });

  it("should mark the phrase as started and store the public key", async function () {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);
    const phrase = await aaa.getPhrase(pid);
    expect(phrase.started).to.be.true;
    expect(phrase.pk).to.equal(ethers.hexlify(pk));
  });

  it("should choose n nodes and produce n WordRequested events", async function () {
    const pid = generateRandomBytes(32);
    const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);
    const receipt = await tx.wait();
    const parsedLogs = receipt.logs
      .map((log: any) => {
        try {
          return aaa.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .filter((log: { name: string }) => log?.name === "WordRequested");
    expect(parsedLogs.length).to.equal(WORDS_NEEDED);
  });

  it("should emit WordRequested events before SeedPhraseProtocolInitiated", async function () {
    const pid = generateRandomBytes(32);
    const tx = await aaa.seedPhraseGenerationProtocol(pid, pk);
    const receipt = await tx.wait();

    const parsed = receipt.logs
      .map((log: any) => {
        try {
          return aaa.interface.parseLog(log);
        } catch {
          return null;
        }
      })
      .filter((log: any) => log);

    const wordEvents = parsed.filter((e: any) => e.name === "WordRequested");
    const seedEvent = parsed.find(
      (e: any) => e.name === "SeedPhraseProtocolInitiated"
    );

    expect(wordEvents.length).to.equal(WORDS_NEEDED);
    const seedIndex = parsed.indexOf(seedEvent);
    const lastWordIndex = parsed.lastIndexOf(wordEvents[wordEvents.length - 1]);
    expect(lastWordIndex).to.be.lessThan(seedIndex);
  });
});
