package methods

import (
	"go-structure-diagram/pkg/diagram"
	"go/ast"
)

// Analyze extracts interactions related to method and function calls.
func Analyze(file *ast.File, projectPath string) []diagram.Interaction {
	var interactions []diagram.Interaction
	packageName := file.Name.Name

	// Maps to store information about structs, interfaces, and functions
	structs := make(map[string]bool)
	interfaces := make(map[string]bool)
	functions := make(map[string]bool)

	// Identify structs, interfaces, and functions
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					switch s.Type.(type) {
					case *ast.StructType:
						structs[s.Name.Name] = true
					case *ast.InterfaceType:
						interfaces[s.Name.Name] = true
					}
				}
			}
		case *ast.FuncDecl:
			funcName := d.Name.Name
			functions[funcName] = true
		}
	}

	// Use context to track the current function being analyzed
	var currentFunc string

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			currentFunc = node.Name.Name
			defer func() { currentFunc = "" }()
		case *ast.CallExpr:
			// Extract function calls
			if fun, ok := node.Fun.(*ast.Ident); ok {
				if functions[fun.Name] {
					interactions = append(interactions, diagram.Interaction{
						Type:    "func_call",
						Package: packageName,
						From:    currentFunc,
						To:      fun.Name,
						Message: "calls",
					})
				}
			}
			// Extract method calls
			if sel, ok := node.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := sel.X.(*ast.Ident); ok {
					interactions = append(interactions, diagram.Interaction{
						Type:    "method_call",
						Package: packageName,
						From:    currentFunc,
						To:      ident.Name + "." + sel.Sel.Name,
						Message: "calls",
					})
				}
			}
		}
		return true
	})

	return interactions
}
