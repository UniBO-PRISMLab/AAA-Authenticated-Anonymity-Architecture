import { ethers } from "ethers";
import { encryptWithKey, generatePublicKey } from "../utils/crypto";

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
  "function submitEncryptedWord(bytes32 pid, bytes encryptedWord, bytes nodePK) external",
  "event PhraseComplete(bytes32 indexed pid)",
  "event SIDEncryptionRequested(bytes32 indexed pid, address indexed node, bytes sid, bytes userPK)",
  "function storeEncryptedSID(bytes32 pid, bytes encryptedSID) external",
];

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
      const encryptedWord = ethers.hexlify(
        ethers.toUtf8Bytes("enc_" + randomWord)
      );

      const nodePk = await generatePublicKey();
      const nodePkBytes = ethers.toUtf8Bytes(nodePk.toString("base64"));

      try {
        const tx = await contract.submitEncryptedWord(
          pid,
          encryptedWord,
          nodePkBytes,
          {
            gasLimit: 5_000_000,
          }
        );
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

  contract.on(
    "SIDEncryptionRequested",
    async (pid: string, nodes: string, sid: string, userPK: string) => {
      const { buffer, pem } = recoverPublicKey(userPK);

      if (nodes.toLowerCase() !== wallet.address.toLowerCase()) return;

      const encSID = await encryptWithKey(sid, buffer);
      contract.storeEncryptedSID(
        pid,
        ethers.toUtf8Bytes(encSID.toString("base64")),
        {
          gasLimit: 5_000_000,
        }
      );
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

function recoverPublicKey(userPK: string): {
  buffer: Buffer;
  pem: string;
} {
  const hex = userPK.startsWith("0x") ? userPK.slice(2) : userPK;
  const base64 = Buffer.from(hex, "hex").toString("utf-8");
  const buffer = Buffer.from(base64, "base64");
  const pemBody = base64.match(/.{1,64}/g)?.join("\n") ?? base64;
  const pem = `-----BEGIN RSA PUBLIC KEY-----\n${pemBody}\n-----END RSA PUBLIC KEY-----`;

  return { buffer, pem };
}

main().catch(console.error);
