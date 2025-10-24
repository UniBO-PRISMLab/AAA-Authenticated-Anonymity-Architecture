import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { generatePublicKey, generateRandomBytes } from "../utils/crypto";
import { createPublicKey, generateKeyPairSync } from "crypto";

const abi: string[] = [
  "function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey)",
  "event WordRequested(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "event SeedPhraseProtocoloInitiated(bytes32 indexed pid)",
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
  const pid: Buffer = generateRandomBytes(32);
  const { publicKey, privateKey } = generateKeyPairSync("rsa", {
    modulusLength: 2048,
    publicKeyEncoding: { type: "spki", format: "pem" },
    privateKeyEncoding: { type: "pkcs8", format: "pem" },
  });
  console.log("Submitting PID (b64): ", pid.toString("base64"));
  console.log("Submitting PK: ", publicKey);

  const tx: ethers.TransactionResponse =
    await contract.seedPhraseGenerationProtocol(pid, Buffer.from(publicKey));
  console.log("Tx hash:", tx.hash);
}

main().catch((error: any) => {
  console.error("\nError:", error.message);
  console.error(error);
});
