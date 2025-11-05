import { ethers } from "ethers";

const provider = new ethers.WebSocketProvider("ws://127.0.0.1:8545");
const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

const privateKeys = [
  // "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
  // "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
  // "0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897",
  // "0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82",
  // "0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
  // "0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
  "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
];

const wordList = ["apple", "banana", "cherry", "dragon", "eagle"];

const abi = [
  "event WordRequested(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "function submitEncryptedWord(bytes32 pid, bytes calldata encryptedWord, bytes calldata nodePK) external",
  "event PhraseComplete(bytes32 indexed pid, bytes encSID)",
  "event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)",
  "function submitEncryptedSID(bytes32 pid, bytes calldata encSID) external",
  "event PIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes32 symK, bytes32 sid)",
  "function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes calldata encPID) external",
];

function startWorker(
  _id: number,
  contract: ethers.Contract,
  wallet: ethers.Wallet
) {
  console.log("UIP Node started at address:", wallet.address);

  contract.on(
    "WordRequested",
    async (_pid: string, _nodes: string, _symK: string, _sid: string) => {
      console.log("WordRequested event received by", wallet.address);
    }
  );

  contract.on(
    "PIDEncryptionRequested",
    async (pid: string, nodes: string, symK: string, sid: string) => {
      console.log("PIDEncryptionRequested event received by", wallet.address);
    }
  );

  contract.on(
    "SIDEncryptionRequested",
    async (pid: string, nodes: string, symK: string, sid: string) => {
      console.log("SIDEncryptionRequested event received by", wallet.address);
    }
  );

  contract.on("PhraseComplete", (pid: string, encSID: string) => {
    console.log("\nPhrase completed for PID:", pid);
    console.log("Encrypted SID:", encSID);
  });
}

async function main() {
  for (let i = 0; i < privateKeys.length; i++) {
    const wallet = new ethers.Wallet(privateKeys[i], provider);
    const contract = new ethers.Contract(contractAddress, abi, wallet);
    startWorker(i, contract, wallet);
  }
}

main().catch(console.error);
