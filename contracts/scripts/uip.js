import { ethers } from "ethers";
import { Buffer } from "buffer";

const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");

const privateKey =
  "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6";
const wallet = new ethers.Wallet(privateKey, provider);

const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

const abi = [
  "event WordRequestedToUIPNode(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external",
];

const contract = new ethers.Contract(contractAddress, abi, wallet);

const wordList = ["apple", "banana", "cherry", "dragon", "eagle"];

// Simple mock "encryption"
function encryptWord(word, publicKey) {
  // In real implementation: ECIES/Hybrid crypto with the provided pubKey
  return ethers.hexlify(ethers.toUtf8Bytes("enc_" + word));
}

async function main() {
  console.log("UIP Node started. Listening for WordRequestedToUIPNode...");

  contract.on("WordRequestedToUIPNode", async (pid, node, publicKey, event) => {
    console.log("Event received:");
    console.log("PID:", pid);
    console.log("Target node:", node);
    console.log("Public key:", publicKey);

    // Only respond if this node was chosen
    if (node.toLowerCase() !== wallet.address.toLowerCase()) {
      console.log("Not for me, ignoring...");
      return;
    }

    console.log("This node was selected! Generating word...");

    // Pick random word
    const randomWord = wordList[Math.floor(Math.random() * wordList.length)];

    // Encrypt with pubKey (mocked)
    const encryptedWord = encryptWord(randomWord, publicKey);

    console.log("Submitting encrypted word:", encryptedWord);

    try {
      const tx = await contract.submitEncryptedWord(pid, encryptedWord);
      await tx.wait();
      console.log("Submitted to chain. Tx hash:", tx.hash);
    } catch (err) {
      console.error("Error submitting word:", err);
    }
  });
}

main().catch(console.error);
