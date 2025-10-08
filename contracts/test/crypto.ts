import Enigma from "@cubbit/enigma";
import * as dotenv from "dotenv";

/** Encrypts a message using the loaded RSA public key.
 * @param msg The message to encrypt
 * @returns a Promise that resolves to the encrypted message
 */
export async function encrypt(msg: string): Promise<Buffer> {
  const rsa = await loadRSA();
  return Enigma.RSA.encrypt(msg, rsa.keypair.public_key);
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
