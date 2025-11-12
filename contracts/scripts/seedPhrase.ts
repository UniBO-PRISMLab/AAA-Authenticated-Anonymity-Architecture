import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { generateRandomBytes } from "../utils/crypto";
import { constants, generateKeyPairSync, privateDecrypt } from "crypto";

const abi: string[] = [
  "function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey)",
  "event WordRequested(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "event SeedPhraseProtocoloInitiated(bytes32 indexed pid)",
  "function getWords(bytes32 pid) external view returns(bytes32[] words)",
  "event PhraseComplete(bytes32 indexed pid, bytes encSID)",
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
  // const pid = generateRandomBytes(32);
  const pid: Buffer = Buffer.from(
    "tkuswtEE463L9z0DWu+op2vfoSpAyEQ3MHgD3WNKkQw=",
    "base64"
  );

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

  console.log("Submitting PID (b64): ", pid.toString("base64"));
  console.log("Submitting PK: ", Buffer.from(publicKey).toString("base64"));
  console.log("Private Key: ", Buffer.from(privateKey).toString("base64"));

  const tx: ethers.TransactionResponse =
    await contract.seedPhraseGenerationProtocol(pid, Buffer.from(publicKey));
  console.log("Tx hash:", tx.hash);

  contract.on("PhraseComplete", (pid: string, encSID: string) => {
    console.log("\nPhrase completed for PID:", pid);
    console.log("Encrypted SID:", encSID);

    const sidBuffer = Buffer.from(encSID.slice(2), "hex");
    const decryptedSID = privateDecrypt(
      {
        key: privateKey,
        padding: constants.RSA_PKCS1_OAEP_PADDING,
        oaepHash: "sha256",
      },
      sidBuffer
    );

    console.log("Decrypted SID (b64):", decryptedSID.toString("base64"));
  });
}

main().catch((error: any) => {
  console.error("\nError:", error.message);
  console.error(error);
});
