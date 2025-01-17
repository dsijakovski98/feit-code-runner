import ts from "typescript";
import { appendPlaceholderComment, isConsoleFunc } from "./utils";

export const cleanUpDebugs = (code: string) => {
  const srcFile = ts.createSourceFile(
    "temp.ts",
    code,
    ts.ScriptTarget.Latest,
    true,
    ts.ScriptKind.TS
  );

  const transformer: ts.TransformerFactory<ts.SourceFile> = (ctx) => {
    return (sourceFile) => {
      const visitor: ts.Visitor = (node): ts.Node | undefined => {
        if (
          ts.isExpressionStatement(node) &&
          ts.isCallExpression(node.expression) &&
          ts.isPropertyAccessExpression(node.expression.expression) &&
          ts.isIdentifier(node.expression.expression.expression) &&
          node.expression.expression.expression.text === "console" &&
          ts.isIdentifier(node.expression.expression.name) &&
          isConsoleFunc(node.expression.expression.name.text)
        ) {
          return undefined;
        }

        return ts.visitEachChild(node, visitor, ctx);
      };

      return ts.visitNode(
        sourceFile,
        (srcFile) => ts.visitEachChild(srcFile, visitor, ctx),
        ts.isSourceFile
      );
    };
  };

  const result = ts.transform(srcFile, [transformer]);
  const printer = ts.createPrinter({ newLine: ts.NewLineKind.LineFeed });
  let cleanCode = printer.printFile(result.transformed[0]);

  return appendPlaceholderComment(cleanCode);
};
