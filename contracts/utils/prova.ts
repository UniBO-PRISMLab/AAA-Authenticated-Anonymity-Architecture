import Enigma from "@cubbit/enigma";
import { loadPublicKey } from "./crypto";

loadPublicKey().then((pk) => {
  console.log("Loaded public key:", pk.toString("base64"));
});
