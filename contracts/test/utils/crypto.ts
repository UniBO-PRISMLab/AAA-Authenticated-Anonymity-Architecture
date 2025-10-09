import Enigma from "@cubbit/enigma";
import * as dotenv from "dotenv";

/**
 * Encrypts a message using the loaded RSA public key.
 * @param msg The message to encrypt
 * @returns a Promise that resolves to the encrypted message
 */
export async function encrypt(msg: string): Promise<Buffer> {
  const rsa = await loadRSA();
  return Enigma.RSA.encrypt(msg, rsa.keypair.public_key);
}

/**
 * Encrypts a message using the provided RSA public key.
 * @param msg The message to encrypt
 * @param publicKey The RSA public key to use for encryption
 * @returns a Promise that resolves to the encrypted message
 */
export async function encryptWithKey(
  msg: string,
  publicKey: Buffer
): Promise<Buffer> {
  return Enigma.RSA.encrypt(msg, publicKey);
}

/**
 * Encrypts a message using the loaded RSA public key.
 * @param msg The encrypted message to decrypt
 * @returns a Promise that resolves to the decrypted message or void
 */
export async function decrypt(msg: Buffer): Promise<string | void> {
  const rsa = await loadRSA();
  const decrypted = (await rsa.decrypt(msg)).toString();
  return decrypted;
}

export async function loadPublicKey(): Promise<Buffer> {
  const keypair = await loadKeypair();
  return keypair.public_key;
}

/**
 * Generates a new RSA public key.
 * @param size The size of the key in bits (default: 2048)
 * @param exponent The public exponent (default: 0x10001)
 * @returns a Promise that resolves to the generated public key as a Buffer
 */
export async function generatePublicKey(
  size: number = 2048,
  exponent: number = 0x10001
): Promise<Buffer> {
  const keypair = await Enigma.RSA.create_keypair({
    size,
    exponent,
  });
  return keypair.public_key;
}

/** Generates random bytes of specified length.
 * @param length The number of random bytes to generate
 * @returns A Buffer containing the random bytes
 */
export function generateRandomBytes(length: number): Buffer {
  return Enigma.Random.bytes(length);
}

/**
 * Loads the RSA keypair from environment variables or generates a new one.
 * @returns a Promise that resolves to an Enigma.RSA.Keypair
 */
function loadKeypair(): Promise<Enigma.RSA.Keypair> {
  dotenv.config({ path: ".env" });

  const public_key_env = process.env.RSA_PUBLIC_KEY;
  const private_key_env = process.env.RSA_PRIVATE_KEY;
  if (public_key_env && private_key_env) {
    const public_key = Buffer.from(public_key_env, "base64");
    const private_key = Buffer.from(private_key_env, "base64");
    const keypair: Enigma.RSA.Keypair = { public_key, private_key };
    return Promise.resolve(keypair);
  }

  const keypair = Enigma.RSA.create_keypair({
    size: 2048,
    exponent: 0x10001,
  });

  return keypair;
}

/**
 * Loads a Enigma.RSA instance.
 * @returns a Promise that resolves to an Enigma.RSA instance
 */
async function loadRSA(): Promise<Enigma.RSA> {
  const keypair = await loadKeypair();
  const rsa = new Enigma.RSA().init({ keypair });
  return rsa;
}
