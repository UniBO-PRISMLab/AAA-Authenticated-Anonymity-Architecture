import "@nomicfoundation/hardhat-toolbox";
import "solidity-docgen";
import type { HardhatUserConfig } from "hardhat/config";

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  gasReporter: {
    enabled: true,
    currency: "EUR",
    reportFormat: "markdown",
    outputFile: "./docs/gas-report.md",
    coinmarketcap: process.env.COINMARKETCAP_API_KEY || "",
    etherscan: process.env.ETHERSCAN_API_KEY || "",
  },
  paths: {
    artifacts: "./artifacts",
    cache: "./cache",
    sources: "./contracts",
    tests: "./test",
  },
  solidity: {
    version: "0.8.28",
    settings: {
      metadata: {
        // Not including the metadata hash
        // https://github.com/paulrberg/hardhat-template/issues/31
        bytecodeHash: "none",
      },
      // Disable the optimizer when debugging
      // https://hardhat.org/hardhat-network/#solidity-optimizer-support
      optimizer: {
        enabled: true,
        runs: 800,
      },
    },
  },
  typechain: {
    outDir: "types",
    target: "ethers-v6",
  },
  docgen: {
    outputDir: "./docs",
    pages: "files",
  },
};

export default config;
