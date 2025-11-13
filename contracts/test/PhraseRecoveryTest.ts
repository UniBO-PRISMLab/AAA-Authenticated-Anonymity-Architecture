import { expect } from "chai";
import { ethers } from "hardhat";
import { getRandomWord } from "../utils/dictionary";
import {
  generateKeypair,
  encryptWithKey,
  decryptWithKey,
  generateRandomBytes,
} from "../utils/crypto";

describe("User Phrase Recovery", function () {
  const WORDS_NEEDED = 4;
  const REDUNDANCY_FACTOR = 2;

  let owner: any;
  let nodes: any[];
  let aaa: any;
  let nodeAddrs: string[];
  let nodeKeypairs: any[];

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    owner = signers[0];
    nodes = signers.slice(1, WORDS_NEEDED + 2);

    const Factory = await ethers.getContractFactory("AAA");
    nodeAddrs = await Promise.all(nodes.map((n) => n.getAddress()));
    nodeKeypairs = await Promise.all(nodes.map(() => generateKeypair()));

    aaa = await Factory.deploy(nodeAddrs, WORDS_NEEDED, REDUNDANCY_FACTOR);
    await aaa.waitForDeployment();
  });

  it("should retrieve the encrypted phrase and decrypt it using user's private key", async function () {
    const pid = generateRandomBytes(32);
    const userKeypair = await generateKeypair();
    const userPk = userKeypair.public_key;

    await aaa.seedPhraseGenerationProtocol(pid, userPk);

    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    const originalWords: string[] = [];

    for (let i = 0; i < selected.length; i++) {
      const nodeAddr = selected[i];
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      const nodeSigner = nodes[nodeIndex];

      const word = await getRandomWord();
      originalWords.push(word);

      const encWord = await encryptWithKey(word, userPk);
      await aaa.connect(nodeSigner).submitEncryptedWord(pid, encWord);
    }

    const [started, pk, encWords] = await aaa.getPhrase(pid);
    expect(started).to.be.true;
    expect(pk).to.equal(ethers.hexlify(userPk));

    const decryptedWords: string[] = [];
    for (const encWord of encWords) {
      let buf: Buffer;
      if (typeof encWord === "string") {
        buf = Buffer.from(encWord.slice(2), "hex");
      } else {
        buf = Buffer.from(encWord);
      }
      const word = await decryptWithKey(buf, userKeypair);
      decryptedWords.push(word);
    }

    console.log("Original Words: ", originalWords);
    console.log("Decrypted Words:", decryptedWords);

    expect(decryptedWords).to.deep.equal(originalWords);
  });

  it("should allow nodes to recover words using redundant submissions", async function () {
    const pid = generateRandomBytes(32);
    const userKeypair = await generateKeypair();
    const userPk = userKeypair.public_key;

    await aaa.seedPhraseGenerationProtocol(pid, userPk);
    const selected = await aaa.connect(owner).getSelectedNodes(pid);
    // Store which redundant node submitted for which index
    const redundantSubmission: number[] = [];
    const originalWords: string[] = [];
    const decRedWords: string[] = [];
    for (let i = 0; i < selected.length; i++) {
      const word = await getRandomWord();
      originalWords.push(word);
      const encWord = await encryptWithKey(word, userPk);

      const nodeAddr = selected[i];
      const nodeIndex = nodeAddrs.findIndex((addr) => addr === nodeAddr);
      const nodeSigner = nodes[nodeIndex];

      await expect(aaa.connect(nodeSigner).submitEncryptedWord(pid, encWord))
        .to.emit(aaa, "WordSubmitted")
        .withArgs(pid, nodeAddr, ethers.keccak256(encWord), i);

      const filter = aaa.filters.RedundantWordRequested(pid, i, null);
      const events = await aaa.queryFilter(filter);
      expect(events.length).to.equal(REDUNDANCY_FACTOR - 1);
      const redundantNodeAddr = events[events.length - 1].args.toNode;
      const redundantNodeIndex = nodeAddrs.findIndex(
        (addr) => addr === redundantNodeAddr
      );
      const redundantNodeSigner = nodes[redundantNodeIndex];
      const nodePk = nodeKeypairs[redundantNodeIndex].public_key;
      const encWordForRedundant = await encryptWithKey(word, nodePk);

      await expect(
        aaa
          .connect(redundantNodeSigner)
          .submitRedundantWord(pid, encWordForRedundant, i, nodePk)
      )
        .to.emit(aaa, "RedundantWordSubmitted")
        .withArgs(
          pid,
          i,
          redundantNodeAddr,
          ethers.keccak256(encWordForRedundant)
        );

      redundantSubmission[i] = redundantNodeIndex;
    }

    for (let i = 0; i < redundantSubmission.length; i++) {
      const [redundantWordsAti, _nodePublicKeys] = await aaa.getRedundantWords(
        pid,
        i
      );

      // Retrieve the redundant node who encrypted this word
      const redundantNodeIndex = redundantSubmission[i];
      const redundantNodeKeyPair = nodeKeypairs[redundantNodeIndex];

      // Decrypt the redundant word and compare with original
      let buf: Buffer;
      const encWord = redundantWordsAti[0];
      if (typeof encWord === "string") {
        buf = Buffer.from(encWord.slice(2), "hex");
      } else {
        buf = Buffer.from(encWord);
      }
      const decryptedRedundantWord = await decryptWithKey(
        buf,
        redundantNodeKeyPair
      );
      expect(decryptedRedundantWord).to.equal(originalWords[i]);
      decRedWords.push(decryptedRedundantWord);
    }

    console.log("Original Words: ", originalWords);
    console.log("Decrypted Redundant Words:", decRedWords);

    expect(decRedWords).to.deep.equal(originalWords);
  });
});
