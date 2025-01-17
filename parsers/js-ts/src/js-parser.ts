import * as acorn from "acorn";
import { generate } from "escodegen";
import * as estraverse from "estraverse"; // Import estraverse
import type { Node } from "estree"; // Import Node type
import { appendPlaceholderComment, isConsoleFunc, type ProcessOptions } from "./utils";

const parseCode = (code: string) => {
  return acorn.parse(code, {
    ecmaVersion: "latest",
    sourceType: "module",
  });
};

export const cleanUpDebugs = (code: string) => {
  const ast = parseCode(code);

  let cleanAst = estraverse.replace(ast as unknown as Node, {
    enter: (node, parent) => {
      const parentExpression = parent?.type === "ExpressionStatement" && parent.expression;

      if (
        parentExpression &&
        node.type === "CallExpression" &&
        node.callee.type === "MemberExpression" &&
        node.callee.object.type === "Identifier" &&
        node.callee.object.name === "console" &&
        node.callee.property.type === "Identifier" &&
        isConsoleFunc(node.callee.property.name)
      ) {
        return estraverse.VisitorOption.Remove;
      }
    },
  });

  cleanAst = estraverse.replace(cleanAst, {
    enter: (node) => {
      if (node.type === "ExpressionStatement" && node.expression === null) {
        return estraverse.VisitorOption.Remove;
      }
    },
  });

  const cleanCode = generate(cleanAst);

  return appendPlaceholderComment(cleanCode);
};
