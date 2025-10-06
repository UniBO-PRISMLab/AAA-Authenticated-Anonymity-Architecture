const { ethers } = require("hardhat");

async function main() {
  console.log("Deploying AAAContract.sol");

  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);

  const nodeAddresses = [
    "0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f",
    "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720",
    "0xBcd4042DE499D14e55001CcbB24a551F3b954096",
    "0x71bE63f3384f5fb98995898A86B02Fb2426c5788",
    "0x2546BcD3c84621e976D8185a91A922aE77ECEc30",
    "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
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
