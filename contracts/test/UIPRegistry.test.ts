import { expect } from "chai";
import { ethers } from "hardhat";

describe("UIPRegistry", function () {
  let owner: any;
  let other: any;
  let node1: any;
  let node2: any;
  let registry: any;

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    node1 = signers[1];
    node2 = signers[2];
    other = signers[3];

    const Factory = await ethers.getContractFactory("MockUIPRegistry");
    registry = await Factory.connect(owner).deploy([await node1.getAddress()]);
    await registry.waitForDeployment();
  });

  describe("constructor", function () {
    it("should set owner and initial nodes", async function () {
      expect(await registry.owner()).to.equal(await owner.getAddress());
      expect(await registry.isNode(await node1.getAddress())).to.be.true;

      const list = await registry.nodeList(0);
      expect(list).to.equal(await node1.getAddress());
    });

    it("should emit NodeAdded for each initial node", async function () {
      const Factory = await ethers.getContractFactory("MockUIPRegistry");
      const tx = await Factory.deploy([
        await node1.getAddress(),
        await node2.getAddress(),
      ]);

      const deploymentTx = tx.deploymentTransaction();
      if (!deploymentTx) {
        throw new Error("Deployment transaction is null");
      }

      const deployed = await tx.waitForDeployment();

      await expect(tx.deploymentTransaction())
        .to.emit(deployed, "NodeAdded")
        .withArgs(await node1.getAddress());
    });
  });

  describe("addNode()", function () {
    it("should allow owner to add new node", async function () {
      const nodeAddr = await node2.getAddress();
      await expect(registry.connect(owner).addNode(nodeAddr))
        .to.emit(registry, "NodeAdded")
        .withArgs(nodeAddr);

      expect(await registry.isNode(nodeAddr)).to.be.true;
      const list = await registry.nodeList(1);
      expect(list).to.equal(nodeAddr);
    });

    it("should revert if non-owner tries to add", async function () {
      const nodeAddr = await node2.getAddress();
      await expect(
        registry.connect(other).addNode(nodeAddr)
      ).to.be.revertedWith("Only owner");
    });

    it("should revert if node already exists", async function () {
      const nodeAddr = await node1.getAddress();
      await expect(
        registry.connect(owner).addNode(nodeAddr)
      ).to.be.revertedWith("exists");
    });
  });

  describe("removeNode()", function () {
    beforeEach(async function () {
      await registry.connect(owner).addNode(await node2.getAddress());
    });

    it("should remove node and emit event", async function () {
      const nodeAddr = await node1.getAddress();
      await expect(registry.connect(owner).removeNode(nodeAddr))
        .to.emit(registry, "NodeRemoved")
        .withArgs(nodeAddr);

      expect(await registry.isNode(nodeAddr)).to.be.false;
    });

    it("should revert if non-owner tries to remove", async function () {
      await expect(
        registry.connect(other).removeNode(await node1.getAddress())
      ).to.be.revertedWith("Only owner");
    });

    it("should revert if node not present", async function () {
      const unknown = ethers.Wallet.createRandom().address;
      await expect(
        registry.connect(owner).removeNode(unknown)
      ).to.be.revertedWith("absent");
    });
  });

  describe("modifiers", function () {
    it("should allow only owner for onlyOwnerFn", async function () {
      await expect(registry.connect(owner).onlyOwnerFn()).to.not.be.reverted;
      await expect(registry.connect(other).onlyOwnerFn()).to.be.revertedWith(
        "Only owner"
      );
    });

    it("should allow only nodes for onlyNodeFn", async function () {
      await expect(registry.connect(node1).onlyNodeFn()).to.not.be.reverted;
      await expect(registry.connect(other).onlyNodeFn()).to.be.revertedWith(
        "Not UIP node"
      );
    });
  });
});
