const { ethers } = require("hardhat");

async function main() {
  console.log("Deploying AAAContract.sol");

  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);

  const nodeAddresses = [
    "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720",
    "0x2222222222222222222222222222222222222222",
    "0x3333333333333333333333333333333333333333",
    "0x4444444444444444444444444444444444444444",
    "0x5555555555555555555555555555555555555555",
    "0x6666666666666666666666666666666666666666",
    "0x7777777777777777777777777777777777777777",
    "0x8888888888888888888888888888888888888888",
    "0x9999999999999999999999999999999999999999",
    "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
    "0xcccccccccccccccccccccccccccccccccccccccc",
  ];

  const AAAContract = await ethers.getContractFactory("AAAContract");
  const aaaContract = await AAAContract.deploy(nodeAddresses);
  await aaaContract.waitForDeployment();

  console.log("AAAContract deployed to:", await aaaContract.getAddress());
  console.log(`Initialized with ${nodeAddresses.length} nodes`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
