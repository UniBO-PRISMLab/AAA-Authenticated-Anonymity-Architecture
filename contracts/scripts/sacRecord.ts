import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { generateKeyPairSync } from "crypto";

const abi: string[] = [
  "function submitSACRecord(bytes sac, bytes pk)",
  "function getSACRecord(bytes pk) external view returns (bytes)",
];

const provider: ethers.JsonRpcProvider = new ethers.JsonRpcProvider(
  "http://127.0.0.1:8545"
);

const privateKey: string =
  "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97";
const wallet: Wallet = new ethers.Wallet(privateKey, provider);
const contractAddress: string = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
const contract: Contract = new ethers.Contract(contractAddress, abi, wallet);

async function main(): Promise<void> {
  const sac = "7MWorstwntw=";

  const { publicKey, privateKey } = generateKeyPairSync("rsa", {
    modulusLength: 2048,
    publicExponent: 0x10001,
    publicKeyEncoding: {
      type: "spki",
      format: "pem",
    },
    privateKeyEncoding: {
      type: "pkcs8",
      format: "pem",
    },
  });

  const sacBytes = Buffer.from(sac, "base64");
  const pkBytes = Buffer.from(publicKey);

  console.log("Submitting SAC Record...");
  console.log("SAC:", sac);
  console.log(
    "Public Key (Base64):",
    Buffer.from(publicKey).toString("base64")
  );

  const tx = await contract.submitSACRecord(sacBytes, pkBytes);
  console.log("Tx sent:", tx.hash);

  const receipt = await tx.wait();
  console.log("Transaction mined:", receipt.transactionHash);

  const storedSAC = await contract.getSACRecord(pkBytes);
  console.log(
    "Retrieved SAC from blockchain:",
    Buffer.from(storedSAC.slice(2), "hex").toString("base64")
  );
}

main().catch((err) => {
  console.error("Error:", err);
});
