// import { ethers } from "ethers";

// const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");
// const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
// const privateKeys = [
//   "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
//   "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
//   "0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897",
//   "0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82",
//   "0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
//   "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e",
// ];

// const abi = [
//   "event WordRequestedToUIPNode(bytes32 indexed pid, address indexed node, bytes publicKey)",
//   "function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external",
//   "event PhraseComplete(bytes32 indexed pid)",
// ];

// const wordList = ["apple", "banana", "cherry", "dragon", "eagle"];

// function encryptWord(word, publicKey) {
//   // In real implementation: ECIES/Hybrid crypto with the provided pubKey
//   return ethers.hexlify(ethers.toUtf8Bytes("enc_" + word));
// }

// function startWorker(id, contract, wallet) {
//   console.log("UIP Node started at address:", wallet.address);

//   contract.on("WordRequestedToUIPNode", async (pid, node, publicKey, event) => {
//     console.log("\nEvent received:");
//     console.log("PID:", pid);
//     console.log("Target node:", node);
//     console.log("Public key:", publicKey);

//     if (node.toLowerCase() !== wallet.address.toLowerCase()) {
//       console.log("Not for me, ignoring...");
//       return;
//     }

//     console.log("This node was selected! Generating word...");

//     const randomWord = wordList[Math.floor(Math.random() * wordList.length)];

//     const encryptedWord = encryptWord(randomWord, publicKey);

//     console.log("Submitting encrypted word:", encryptedWord);

//     try {
//       const tx = await contract.submitEncryptedWord(pid, encryptedWord, {
//         gasLimit: 5_000_000,
//       });
//       await tx.wait();
//       console.log("Submitted to chain. Tx hash:", tx.hash);
//     } catch (err) {
//       console.error("Error submitting word:", err);
//     }
//   });

//   contract.on("PhraseComplete", (pid, phrase, event) => {
//     console.log(`Phrase completed for PID ${pid}: ${phrase}`);
//   });
// }

// async function main() {
//   for (let i = 0; i < privateKeys.length; i++) {
//     let wallet = new ethers.Wallet(privateKeys[i], provider);
//     let contract = new ethers.Contract(contractAddress, abi, wallet);
//     startWorker(i, contract, wallet);
//   }
// }

// main().catch(console.error);

import { ethers } from "ethers";

const provider = new ethers.JsonRpcProvider("http://127.0.0.1:8545");
const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

const privateKeys = [
  "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
  "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
  "0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897",
  "0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82",
  "0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
  "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e",
];

const wordList = ["apple", "banana", "cherry", "dragon", "eagle"];

const abi = [
  "event WordRequestedToUIPNode(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external",
  "event PhraseComplete(bytes32 indexed pid)",
];

function encryptWord(word: string, publicKey: string): string {
  return ethers.hexlify(ethers.toUtf8Bytes("enc_" + word));
}

function startWorker(
  id: number,
  contract: ethers.Contract,
  wallet: ethers.Wallet
) {
  console.log("UIP Node started at address:", wallet.address);
  contract.on(
    "WordRequestedToUIPNode",
    async (pid: string, node: string, publicKey: string) => {
      if (node.toLowerCase() !== wallet.address.toLowerCase()) return;

      const randomWord = wordList[Math.floor(Math.random() * wordList.length)];
      const encryptedWord = encryptWord(randomWord, publicKey);

      try {
        const tx = await contract.submitEncryptedWord(pid, encryptedWord, {
          gasLimit: 5_000_000,
        });
        await tx.wait();
        console.log(
          `\nNode ${wallet.address} submitted word for PID ${pid}. Tx hash:`,
          tx.hash
        );
      } catch (err: any) {
        console.error(err?.error?.message || err);
      }
    }
  );

  contract.on("PhraseComplete", (pid: string) => {
    console.log("\nPhrase completed for PID:", pid);
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
