import { expect } from "chai";
import { ethers } from "hardhat";

describe("AAAContract", function () {
  const WORDS_NEEDED = 4;
  const REDUNDANCY_M = 2;

  let owner: any;
  let nodes: any[];
  let aaa: any;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 1);

    const Factory = await ethers.getContractFactory("AAAContract");
    const nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_M);
    await aaa.waitForDeployment();
  });

  describe("Deployment", function () {
    it("should deploy with correct parameters", async function () {
      expect(await aaa.WORDS()).to.equal(WORDS_NEEDED);
      expect(await aaa.REDUNDANCY_FACTOR()).to.equal(REDUNDANCY_M);
      expect(await aaa.owner()).to.equal(await owner.getAddress());
    });

    it("should revert if deployed with too few nodes", async function () {
      const Factory = await ethers.getContractFactory("AAAContract");
      const oneNode = [await nodes[0].getAddress()];
      await expect(Factory.deploy(oneNode, 2, REDUNDANCY_M)).to.be.revertedWith(
        "too few nodes"
      );
    });
  });
});
