import { ethers, Contract, Wallet } from "ethers";
import { Buffer } from "buffer";
import { constants, generateKeyPairSync, privateDecrypt } from "crypto";
import fs from "fs";
import os from "os";
import path from "path";

const RPC_URL = "http://127.0.0.1:8545";
const PRIVATE_KEY =
  "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97";
const CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

const ABI = [
  "function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey)",
  "event PhraseComplete(bytes32 indexed pid, bytes encSID)",
];

const PUBLIC_IDS_FILE = path.join(os.homedir(), "public-identities.txt");
const OUTPUT_FILE = path.join(os.homedir(), "anonymous-identities.txt");
const STUDENT_FILE = path.join(
  os.homedir(),
  "student-anonymous-identities.txt"
);

const provider = new ethers.JsonRpcProvider(RPC_URL);
const wallet = new Wallet(PRIVATE_KEY, provider);
const contract = new Contract(CONTRACT_ADDRESS, ABI, wallet);

function loadPIDs(): Buffer[] {
  const lines = fs
    .readFileSync(PUBLIC_IDS_FILE, "utf8")
    .split("\n")
    .map((l) => l.trim())
    .filter((l) => l.length > 0)
    .map((l) => {
      const parts = l.split(":");
      if (parts.length < 2) throw new Error(`Invalid line: ${l}`);
      return parts[1].trim();
    });

  return lines.map((b64) => Buffer.from(b64, "base64"));
}

function getRandomIndices(total: number, count: number): number[] {
  const indices = new Set<number>();
  while (indices.size < count) {
    indices.add(Math.floor(Math.random() * total));
  }
  return [...indices];
}

async function runForPID(pid: Buffer): Promise<{
  pid: string;
  decryptedSID: string;
  publicKey: string;
  privateKey: string;
}> {
  console.log("\nRunning protocol for PID:", pid.toString("base64"));
  const { publicKey, privateKey } = generateKeyPairSync("rsa", {
    modulusLength: 2048,
    publicExponent: 0x10001,
    publicKeyEncoding: { type: "spki", format: "pem" },
    privateKeyEncoding: { type: "pkcs8", format: "pem" },
  });

  const tx = await contract.seedPhraseGenerationProtocol(
    pid,
    Buffer.from(publicKey)
  );
  await tx.wait();
  console.log("Tx:", tx.hash);

  const sid: string = await new Promise((resolve) => {
    const listener = (evPid: string, encSID: string) => {
      const eventPid = Buffer.from(evPid.slice(2), "hex").toString("base64");
      if (eventPid === pid.toString("base64")) {
        contract.off("PhraseComplete", listener);
        resolve(encSID);
      }
    };

    contract.on("PhraseComplete", listener);
  });

  console.log("Encrypted SID (hex):", sid);

  const sidBuffer = Buffer.from(sid.slice(2), "hex");
  const decrypted = privateDecrypt(
    {
      key: privateKey,
      padding: constants.RSA_PKCS1_OAEP_PADDING,
      oaepHash: "sha256",
    },
    sidBuffer
  );

  const decoded = decrypted.toString("base64");
  console.log("Decrypted SID (b64):", decoded);

  return {
    pid: pid.toString("base64"),
    decryptedSID: decrypted.toString("base64"),
    publicKey,
    privateKey,
  };
}

async function main() {
  const pids = loadPIDs();
  if (pids.length === 0) {
    console.error("No PIDs found in", PUBLIC_IDS_FILE);
    return;
  }

  const indices = getRandomIndices(pids.length, 50);
  console.log("Selected random indices:", indices);

  let count = 0;
  const sids = [];
  for (const idx of indices) {
    const pid = pids[idx];
    try {
      const {
        pid: pidStr,
        decryptedSID,
        publicKey,
        privateKey,
      } = await runForPID(pid);
      sids.push(decryptedSID);
      const output = `${count}: ${pidStr}\nSID: ${decryptedSID}\nPublic Key: ${Buffer.from(
        publicKey
      ).toString("base64")}\nPrivate Key: ${Buffer.from(privateKey).toString(
        "base64"
      )}\n\n`;

      fs.appendFileSync(OUTPUT_FILE, output);
      count++;
    } catch (err) {
      console.error("Error for PID: ", pid.toString("base64"), err);
    }
    await new Promise((res) => setTimeout(res, 2000));
  }

  const sidIndices = getRandomIndices(sids.length, 3);
  let cc = 1;
  for (const sidIdx of sidIndices) {
    fs.appendFileSync(STUDENT_FILE, `${cc}: ${sids[sidIdx]}\n`);
    cc++;
  }
  console.log("\nOutput saved to:", OUTPUT_FILE);
  console.log("\nGenerated challenge to:", STUDENT_FILE);
}

main().catch((err) => {
  console.error("Error:", err);
});
