export type ProcessOptions = {
  code: string;
  funcName: string;
  args: unknown[];
};

const consolePrintMethods = ["log", "info", "warn", "error", "debug", "trace", "table", "assert"];

export const isConsoleFunc = (funcName: string) => {
  return consolePrintMethods.includes(funcName);
};

export const appendPlaceholderComment = (cleanCode: string) => {
  cleanCode += "\n";
  cleanCode += "// PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER";

  return cleanCode;
};
