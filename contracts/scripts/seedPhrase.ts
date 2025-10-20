import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { generatePublicKey } from "../utils/crypto";

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
  const pidBase64: string = "Nifl3y+2jnuAxF26jqpjogu0ZYnA6IxZikjmTnnzm1k=";
  const pidBytes: Buffer = Buffer.from(pidBase64, "base64");
  const pidHex: string = "0x" + pidBytes.toString("hex");

  const pubKey: Buffer = await generatePublicKey();

  console.log("\nSubmitting PID...");
  const pkBytes = ethers.toUtf8Bytes(pubKey.toString("base64"));

  const tx: ethers.TransactionResponse =
    await contract.seedPhraseGenerationProtocol(pidHex, pkBytes);
  console.log("Tx hash:", tx.hash);
}

main().catch((error: any) => {
  console.error("\nError:", error.message);
  console.error(error);
});
