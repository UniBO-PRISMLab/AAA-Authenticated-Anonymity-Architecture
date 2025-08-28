import { expect } from "chai";
import { ethers } from "hardhat";
import { AAASmartContract, AAASmartContract__factory } from "../types";

describe("AAASmartContract", function () {
  let contract: AAASmartContract;
  let owner: any;
  let user: any;
  let uipNode: any;
  let anotherUipNode: any;

  const PID = ethers.encodeBytes32String("user1");
  const USER_PK = "UserPublicKey";
  const SAC = ethers.encodeBytes32String("SAC123");
  const ANON_PK = "AnonymousPK";

  beforeEach(async function () {
    [owner, user, uipNode, anotherUipNode] = await ethers.getSigners();

    const factory: AAASmartContract__factory = (await ethers.getContractFactory(
      "AAASmartContract"
    )) as AAASmartContract__factory;

    contract = await factory.deploy();
    await contract.waitForDeployment();

    // Register UIP nodes for submission permissions
    await contract.registerUIP(uipNode.address, true);
    await contract.registerUIP(anotherUipNode.address, true);
  });

  it("deploys and sets owner", async function () {
    expect(await contract.owner()).to.equal(owner.address);
  });

  it("allows seed phrase protocol initiation", async function () {
    await expect(contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK))
      .to.emit(contract, "SeedPhraseProtocolInitiated")
      .withArgs(PID, USER_PK);

    expect(await contract.pidToUserPublicKey(PID)).to.equal(USER_PK);
    expect(await contract.seedPhraseProtocolInitiated(PID)).to.equal(true);
  });

  it("prevents initiating protocol twice for the same PID", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);
    await expect(
      contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK)
    ).to.be.revertedWith(
      "AAAC: Seed phrase protocol already initiated for this PID."
    );
  });

  it("rejects non-UIP word fragment submissions", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);
    await expect(
      contract.connect(user).submitEncryptedWordFragment(PID, 1, "word1")
    ).to.be.revertedWith(
      "AAAC: Only registered UIP nodes can submit word fragments."
    );
  });

  it("rejects invalid word number and empty word", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);

    await expect(
      contract.connect(uipNode).submitEncryptedWordFragment(PID, 0, "w")
    ).to.be.revertedWith("AAAC: Invalid word number.");

    await expect(
      contract.connect(uipNode).submitEncryptedWordFragment(PID, 25, "w")
    ).to.be.revertedWith("AAAC: Invalid word number.");

    await expect(
      contract.connect(uipNode).submitEncryptedWordFragment(PID, 1, "")
    ).to.be.revertedWith("AAAC: Encrypted word cannot be empty.");
  });

  it("submits 24 fragments, finalizes SID, and stores encrypted SID", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);

    // submit all 24 words
    for (let i = 1; i <= 24; i++) {
      const tx = await contract
        .connect(uipNode)
        .submitEncryptedWordFragment(PID, i, `word${i}`);

      await expect(tx)
        .to.emit(contract, "EncryptedWordFragmentSubmitted")
        .withArgs(PID, i, uipNode.address);

      const storedWord = await contract.encryptedWordsForUser(PID, i);
      expect(storedWord).to.equal(`word${i}`);
    }

    expect(await contract.wordsReceivedCount(PID)).to.equal(24);

    const sid = await contract.pidToSid(PID);
    expect(sid).to.not.equal(ethers.ZeroHash);

    const encryptedSid = await contract.encryptedSidForPid(PID);
    expect(encryptedSid).to.include(USER_PK);
  });

  it("allows redundant encrypted word submission to another UIP node", async function () {
    await expect(
      contract
        .connect(uipNode)
        .submitRedundantEncryptedWord(
          PID,
          1,
          "redundantWord",
          anotherUipNode.address
        )
    ).to.not.be.reverted;

    const storedRedundant = await contract.redundantEncryptedWords(
      PID,
      1,
      anotherUipNode.address
    );
    expect(storedRedundant).to.equal("redundantWord");
  });

  it("registers SAC association and verifies it", async function () {
    await expect(contract.registerSACAssociation(SAC, ANON_PK))
      .to.emit(contract, "SacRegistered")
      .withArgs(SAC, ANON_PK);

    expect(
      await contract.checkSACExistenceAndAssociation(SAC, ANON_PK)
    ).to.equal(true);

    expect(
      await contract.checkSACExistenceAndAssociation(SAC, "wrongPK")
    ).to.equal(false);

    // duplicate registration should revert
    await expect(
      contract.registerSACAssociation(SAC, "AnotherPK")
    ).to.be.revertedWith(
      "AAAC: SAC already registered with an anonymous account."
    );
  });

  it("retrieves SYMK-encrypted PID string from SID after finalization", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);
    for (let i = 1; i <= 24; i++) {
      await contract
        .connect(uipNode)
        .submitEncryptedWordFragment(PID, i, `word${i}`);
    }

    const sid = await contract.pidToSid(PID);
    // Should be the ASCII-safe string we built after the contract fix
    const retrieved = await contract.getPIDFromSID(sid);
    expect(retrieved).to.match(/^ENC_PID_WITH_SYMK_FOR_0x[0-9a-fA-F]{64}$/);
  });

  it("returns encrypted seed phrase, encrypted SID, SYMK association, and user PK", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);
    for (let i = 1; i <= 24; i++) {
      await contract
        .connect(uipNode)
        .submitEncryptedWordFragment(PID, i, `word${i}`);
    }

    const [encryptedWords, encryptedSID, symkEncryptedPid, pkUser] =
      await contract.getEncryptedSeedPhraseAndSID(PID);

    expect(encryptedWords.length).to.equal(24);
    expect(encryptedWords[0]).to.equal("word1");
    expect(encryptedWords[23]).to.equal("word24");

    expect(encryptedSID).to.include("ENC_SID_WITH_");
    expect(pkUser).to.equal(USER_PK);
    expect(symkEncryptedPid).to.match(
      /^ENC_PID_WITH_SYMK_FOR_0x[0-9a-fA-F]{64}$/
    );
  });

  it("reverts getEncryptedSeedPhraseAndSID if protocol not initiated", async function () {
    await expect(contract.getEncryptedSeedPhraseAndSID(PID)).to.be.revertedWith(
      "AAAC: Seed phrase protocol not initiated for this PID."
    );
  });

  it("reverts getEncryptedSeedPhraseAndSID if SID not created yet", async function () {
    await contract.connect(user).requestSeedPhraseProtocol(PID, USER_PK);
    await expect(contract.getEncryptedSeedPhraseAndSID(PID)).to.be.revertedWith(
      "AAAC: SID not yet created for this PID. Please wait for protocol completion."
    );
  });

  it("reverts getPIDFromSID if association missing", async function () {
    // random SID (not created by the contract)
    const randomSid = ethers.keccak256(ethers.toUtf8Bytes("random"));
    await expect(contract.getPIDFromSID(randomSid)).to.be.revertedWith(
      "AAAC: PID-SYMK association not found for this SID."
    );
  });
});
