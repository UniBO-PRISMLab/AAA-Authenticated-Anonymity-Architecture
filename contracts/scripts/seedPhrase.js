import { ethers } from "ethers";
import { Buffer } from "buffer";

const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");

const privateKey =
  "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97";
const wallet = new ethers.Wallet(privateKey, provider);

const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

const abi = [
  "function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey)",
  "event WordRequestedToUIPNode(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "event SeedPhraseProtocoloInitiated(bytes32 indexed pid)",
];

const contract = new ethers.Contract(contractAddress, abi, wallet);

async function main() {
  const pidBase64 = "Nifl3y+2jmuAxF26jqpjogu0ZYnA6IxSikjmTnnjm7k=";
  const pidBytes = Buffer.from(pidBase64, "base64");
  const pidHex = "0x" + pidBytes.toString("hex");

  const pubKey =
    "0x04abcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab";

  console.log("\nSubmitting PID...");
  console.log("PID (hex):", pidHex);
  console.log("Public Key:", pubKey);

  const tx = await contract.seedPhraseGenerationProtocol(pidHex, pubKey);
  console.log("Tx hash:", tx.hash);

  const receipt = await tx.wait();
  console.log("Mined:", receipt.transactionHash);

  for (const log of receipt.logs) {
    try {
      const parsed = contract.interface.parseLog(log);
      console.log("Event:", parsed.name, parsed.args);
    } catch {}
  }
}

main().catch((error) => {
  console.error("\nError:", error.message);
  console.error(error);
});
