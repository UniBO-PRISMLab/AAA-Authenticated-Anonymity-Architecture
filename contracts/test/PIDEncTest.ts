import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { ethers } from "hardhat";
import { getRandomWord } from "../utils/dictionary";
import {
  generatePublicKey,
  generateRandomBytes,
  encryptWithKey,
} from "../utils/crypto";

describe("PID Encryption Submission", function () {
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

  async function prepareFullPhrase() {
    const pid = generateRandomBytes(32);
    await aaa.seedPhraseGenerationProtocol(pid, pk);

    const selected = await aaa.connect(owner).getSelectedNodes(pid);

    for (const nodeAddr of selected) {
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      const nodeSigner = nodes[nodeIndex];
      const word = await getRandomWord().then((w) => encryptWithKey(w, pk));
      await aaa.connect(nodeSigner).submitEncryptedWord(pid, word);
    }

    const events = await aaa.queryFilter(
      aaa.filters.SIDEncryptionRequested(pid)
    );
    const encryptionResp = events[0].args.node;
    const encSid = ethers.randomBytes(32);

    const respIndex = nodeAddrs.findIndex((addr) => addr === encryptionResp);
    const respSigner = nodes[respIndex];

    await aaa.connect(respSigner).submitEncryptedSID(pid, encSid);
    return { pid, sid: ethers.keccak256(encSid), respSigner };
  }

  it("should revert if called by non-UIP node", async function () {
    const { pid, sid } = await prepareFullPhrase();
    const encPID = ethers.randomBytes(32);

    await expect(
      aaa.connect(nonNode).submitEncryptedPID(pid, sid, encPID)
    ).to.be.revertedWith("Not UIP node");
  });

  it("should store encrypted PID and emit PhraseComplete", async function () {
    const { pid, sid, respSigner } = await prepareFullPhrase();
    const encPID = ethers.randomBytes(64);

    await expect(aaa.connect(respSigner).submitEncryptedPID(pid, sid, encPID))
      .to.emit(aaa, "PhraseComplete")
      .withArgs(pid, anyValue);

    // Verify the SIDRecord was stored
    const [storedEncPID, storedPK] = await aaa.getSIDRecord(sid);
    expect(storedEncPID).to.equal(ethers.hexlify(encPID));
    expect(storedPK).to.equal(ethers.hexlify(pk));
  });

  it("should revert if already stored", async function () {
    const { pid, sid, respSigner } = await prepareFullPhrase();
    const encPID = ethers.randomBytes(64);

    await aaa.connect(respSigner).submitEncryptedPID(pid, sid, encPID);

    // Try again with the same SID
    await expect(
      aaa.connect(respSigner).submitEncryptedPID(pid, sid, encPID)
    ).to.be.revertedWith("already stored");
  });
});
