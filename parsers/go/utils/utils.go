package utils

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

func createAstFile(code string) (*token.FileSet, *ast.File, error) {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "src.go", code, 0)
	if err != nil {
		return nil, nil, err
	}

	return fileSet, astFile, nil
}

func findStatement(node ast.Node, decls []ast.Decl) (ast.Node, bool) {
	for _, decl := range decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			for _, stmt := range funcDecl.Body.List {
				if containsNode(stmt, node) {
					return stmt, true
				}
			}
		}
	}
	return nil, false
}

func containsNode(stmt ast.Node, node ast.Node) bool {
	var found bool
	ast.Inspect(stmt, func(n ast.Node) bool {
		if n == node {
			found = true
			return false // Stop traversal once found
		}
		return true
	})
	return found
}

func removeStatement(stmt ast.Node, astFile *ast.File) {
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		for j, s := range funcDecl.Body.List {
			if s == stmt {
				funcDecl.Body.List = append(funcDecl.Body.List[:j], funcDecl.Body.List[j+1:]...)
				return
			}
		}
	}
}

func generateCode(fileSet *token.FileSet, astFile *ast.File) (string, error) {
	var buf strings.Builder

	err := format.Node(&buf, fileSet, astFile)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func addFmtImport(astFile *ast.File) {
	// Add 'fmt' import if not there
	importExists := false
	for _, imp := range astFile.Imports {
		fmt.Println(imp.Path.Value)
		if imp.Path.Value == `"fmt"` {
			importExists = true
			break
		}
	}

	if importExists {
		return
	}

	newImport := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"fmt"`,
		},
	}

	astFile.Imports = append(astFile.Imports, newImport)
	astFile.Decls = append([]ast.Decl{
		&ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{newImport},
		},
	}, astFile.Decls...)
}

func CleanupDebugs(code string) (string, error) {
	fileSet, astFile, err := createAstFile(code)
	if err != nil {
		return "", fmt.Errorf("parsing error: %w", err)
	}

	var nodesToRemove []ast.Node // Collect nodes to remove

	// Remove Print statements
	ast.Inspect(astFile, func(node ast.Node) bool {
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		ident, ok := selExpr.X.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Name == "fmt" || ident.Name == "log" {
			if strings.HasPrefix(selExpr.Sel.Name, "Print") {
				if stmt, ok := findStatement(node, astFile.Decls); ok {
					nodesToRemove = append(nodesToRemove, stmt)
				}
			}
		}

		return true
	})

	for _, nodeToRemove := range nodesToRemove {
		removeStatement(nodeToRemove, astFile)
	}

	resultCode, err := generateCode(fileSet, astFile)
	if err != nil {
		return "", fmt.Errorf("formatting error: %w", err)
	}

	return resultCode, nil
}

func AppendPlaceholder(code string) (string, error) {
	fileSet, astFile, err := createAstFile(code)
	if err != nil {
		return "", fmt.Errorf("parsing error: %w", err)
	}

	addFmtImport(astFile)

	ast.Inspect(astFile, func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if funcDecl.Name.Name == "main" {
			const commentText = "// PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER"

			funcDecl.Body.List = append(funcDecl.Body.List, &ast.ExprStmt{
				X: &ast.BasicLit{
					Kind:  token.COMMENT,
					Value: commentText,
				},
			})

			return false
		}

		return true
	})

	resultCode, err := generateCode(fileSet, astFile)
	if err != nil {
		return "", fmt.Errorf("formatting error: %w", err)
	}

	return resultCode, nil
}
