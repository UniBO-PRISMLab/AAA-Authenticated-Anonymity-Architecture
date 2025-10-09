import { expect } from "chai";
import { ethers } from "hardhat";
import { getRandomWord } from "./utils/dictionary";

describe("AAALib", function () {
  let harness: any;
  const PID = ethers.keccak256(
    ethers.toUtf8Bytes("Nifl3y+2jmuAxF26jqpjogu0ZYnA6IxSikjmTnnjm7k=")
  );

  beforeEach(async function () {
    const Factory = await ethers.getContractFactory("MockAAALib");
    harness = await Factory.deploy();
    await harness.waitForDeployment();
  });

  describe("selectNodes()", function () {
    it("should select deterministic subset", async function () {
      const pool = Array.from(
        { length: 5 },
        (_, i) => ethers.Wallet.createRandom().address
      );

      const result1 = await harness.testSelectNodes(PID, pool, 3);
      const result2 = await harness.testSelectNodes(PID, pool, 3);

      expect(result1).to.deep.equal(result2);
    });

    it("should revert when pool too small", async function () {
      const pool = [ethers.Wallet.createRandom().address];
      await expect(harness.testSelectNodes(PID, pool, 2)).to.be.revertedWith(
        "pool too small"
      );
    });

    it("should not include duplicates", async function () {
      const pool = Array.from(
        { length: 10 },
        (_, i) => ethers.Wallet.createRandom().address
      );

      const selected = await harness.testSelectNodes(PID, pool, 5);
      const seen = new Set(selected);
      expect(seen.size).to.equal(selected.length);
    });

    it("should return valid addresses from the pool", async function () {
      const pool = Array.from(
        { length: 6 },
        (_, i) => ethers.Wallet.createRandom().address
      );
      const selected = await harness.testSelectNodes(PID, pool, 4);
      selected.forEach((a: string) => expect(pool).to.include(a));
    });
  });

  describe("deriveSymK()", function () {
    it("should return zero for empty array", async function () {
      const result = await harness.testDeriveSymK([]);
      expect(result).to.equal(ethers.ZeroHash);
    });

    it("should be deterministic for same input", async function () {
      const words = [
        await getRandomWord().then(ethers.toUtf8Bytes),
        await getRandomWord().then(ethers.toUtf8Bytes),
      ];
      const r1 = await harness.testDeriveSymK(words);
      const r2 = await harness.testDeriveSymK(words);
      expect(r1).to.equal(r2);
    });

    it("should change output if any word changes", async function () {
      const word1 = await getRandomWord().then(ethers.toUtf8Bytes);
      const base = [word1, await getRandomWord().then(ethers.toUtf8Bytes)];
      const changed = [word1, await getRandomWord().then(ethers.toUtf8Bytes)];
      const r1 = await harness.testDeriveSymK(base);
      const r2 = await harness.testDeriveSymK(changed);
      expect(r1).to.not.equal(r2);
    });

    it("should produce different hashes for different orders", async function () {
      const word1 = await getRandomWord().then(ethers.toUtf8Bytes);
      const word2 = await getRandomWord().then(ethers.toUtf8Bytes);

      const wordsA = [word1, word2];
      const wordsB = [word2, word1];
      const rA = await harness.testDeriveSymK(wordsA);
      const rB = await harness.testDeriveSymK(wordsB);
      expect(rA).to.not.equal(rB);
    });
  });
});
