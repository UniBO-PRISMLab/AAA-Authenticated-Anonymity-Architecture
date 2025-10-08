import * as fs from "fs";
import * as readline from "readline";
import path from "path";
import Enigma from "@cubbit/enigma";

const WORDS_PATH = path.join(__dirname, "../words.txt");

/**
 * Gets a random word from a English word list https://github.com/dwyl/english-words/blob/master/words_alpha.txt.
 * @returns a Promise that resolves to a random word
 */
export async function getRandomWord(): Promise<string> {
  const random_int4 = Enigma.Random.integer(32);
  const word_index = (random_int4 % 370105) + 1;

  const fileStream = fs.createReadStream(WORDS_PATH);
  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  });

  let current = 0;
  for await (const line of rl) {
    current++;
    if (current === word_index) {
      rl.close();
      return line.trim() || `word${word_index}`;
    }
  }

  return `word${word_index}`;
}
