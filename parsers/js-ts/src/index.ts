import { parseArgs } from "node:util";
import { picklist, safeParse } from "valibot";
import { cleanUpDebugs as cleanUpDebugsJs } from "./js-parser";
import { cleanUpDebugs as cleanUpDebugsTs } from "./ts-parser";

const LangSchema = picklist(["js", "ts"]);

const { values } = parseArgs({
  args: Bun.argv,
  options: {
    file: {
      type: "string",
      short: "f",
    },
    lang: {
      type: "string",
      short: "l",
    },
  },
  strict: true,
  allowPositionals: true,
});

if (!values.file) {
  throw new Error("Please provide a file path to parse");
}

if (!values.lang) {
  throw new Error('Please provide a the file\'s language: "js" | "ts"');
}

const langResult = safeParse(LangSchema, values.lang);

if (!langResult.success) {
  throw new Error('Please provide a valid language: "js" | "ts"');
}

const lang = langResult.output;
const filePath = values.file;

let cleanedUpCode = "";
const code = await Bun.file(filePath).text();

if (lang === "js") {
  cleanedUpCode = cleanUpDebugsJs(code);
}

if (lang === "ts") {
  cleanedUpCode = cleanUpDebugsTs(code);
}

await Bun.write(filePath, cleanedUpCode);
