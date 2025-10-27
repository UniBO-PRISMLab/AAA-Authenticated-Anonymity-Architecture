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
  // For deterministic testing
  const pid: Buffer = Buffer.from(
    "wjnuZkNtmbpPl/vROMXgSLfgO+rpZ3FpUGvr1czhT6Q=",
    "base64"
  );
  const publicKey: Buffer = Buffer.from(
    "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFxR2ZvQTJrVEYrZldQWm5kMGZKVwowWGpEMjd1eFVJQnJuUkErQisyc3J0UFZ6R21OVUFqSDNZaVVNVXovZmsySGN6dU5KVFZUVUJzODRaNE41ZDd3ClBhSEhUVjlMeGMwVVhMRWxxMGR1QTdiY0RhaU5TVllGQzNwRjlKVU45c2hGdGxmWXBtd3l3bldDYVBGQlZ4ZE4KUVFvMnNkZGxLRTM5K0hsKy92b0RsUWF5T1c5TE9lb0pzeTdGNWRxYmtwVVVSN0lVUW11UTFFMlpEVG5IcHZkNgo0Y0ZleVltV3pvYlh6WWRjNmZnUjdWS0I5MHRXSXIzRFBMWkRaS2VuOXF0Tkt3QWRDazNrdlVCdW5IYVQxczNyCnRjeGpsejEvd2cvcUxVRWpZQWhaOUlRakRycW92Njg3ZUd0ZlR0Qi9iRVYwV2VMTC9WdkU1UTFkMVJEMk1uNnEKWndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
    "base64"
  );

  // For random testing
  // const pid: Buffer = generateRandomBytes(32);
  // const { publicKey, privateKey } = generateKeyPairSync("rsa", {
  //   modulusLength: 2048,
  //   publicKeyEncoding: { type: "spki", format: "pem" },
  //   privateKeyEncoding: { type: "pkcs8", format: "pem" },
  // });

  console.log("Submitting PID (b64): ", pid.toString("base64"));
  console.log("Submitting PK: ", publicKey.toString("base64"));

  const tx: ethers.TransactionResponse =
    await contract.seedPhraseGenerationProtocol(pid, publicKey);
  console.log("Tx hash:", tx.hash);
}

main().catch((error: any) => {
  console.error("\nError:", error.message);
  console.error(error);
});
