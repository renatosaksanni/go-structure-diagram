package globals

import (
	"go-structure-diagram/pkg/diagram"
	"go/ast"
	"go/token"
)

// Analyze extracts interactions related to global variables and constants.
func Analyze(file *ast.File, projectPath string) []diagram.Interaction {
	var interactions []diagram.Interaction
	packageName := file.Name.Name

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok == token.VAR || d.Tok == token.CONST {
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.ValueSpec:
						for _, name := range s.Names {
							interactions = append(interactions, diagram.Interaction{
								Type:    "global_var",
								Package: packageName,
								Name:    name.Name,
							})
						}
					}
				}
			}
		}
	}

	return interactions
}
