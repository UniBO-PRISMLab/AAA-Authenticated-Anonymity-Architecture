import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { generatePublicKey, generateRandomBytes } from "../utils/crypto";
import { createPublicKey, generateKeyPairSync } from "crypto";

const abi: string[] = [
  "function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey)",
  "event WordRequested(bytes32 indexed pid, address indexed node, bytes publicKey)",
  "event SeedPhraseProtocoloInitiated(bytes32 indexed pid)",
  "function getWords(bytes32 pid) external view returns(bytes32[] words)",
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
    "dy0QlcEMneBt+zTk219tEaTpVYWErxl7umC6pVJmZhE=",
    "base64"
  );
  const publicKey: Buffer = Buffer.from(
    "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFtVmRYZlhXQjBzVGRkanJzWWt0OAp0OXN3elNpUFRiRnhlcnpZUFNKWkkyUlQ3aDBYM3o5VmJkbDN0VTNxaFZJcWFUbHFMeVI5SjBVY1ZnYkpZRVMzCnN5MUF3ejNMUXZBR2hLTzZsczJGTGhPVE5DSUZWb3VxK3cvSnB2b2JDaml5QjFMbjhpNk5HVU55UkEwT3lueTMKU1UwM3owNGM2QlBtYnVrZ3YzL0M1REdVcnNaQ2JmczQ4N0xhYytnakpPbXY5RExjS3VBZWNqdEhNWW54Z1RaVgpUZlROb2h5UDV1Y0tlRExETW5FR250RVQ3enBPenNaWmUrcHJkR0w3T2tEUUpQVkhvYm1xejJzbSs5RFVoRWpOCitieHdRaHd1czRlOVZ5azRPREwydk14cDFFVGVQa2NPbXdiNmJlRTJpUkxkVDVQV2s5YndpMm8wK1FrUzl1QjQKQndJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==",
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

  // const words = await contract.getWords(pid);
  // console.log(words);
}

main().catch((error: any) => {
  console.error("\nError:", error.message);
  console.error(error);
});
