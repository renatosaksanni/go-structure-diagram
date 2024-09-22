package structs

import (
	"go-structure-diagram/pkg/diagram"
	"go/ast"
)

// Analyze extracts interactions related to structs and interfaces.
func Analyze(file *ast.File, projectPath string) []diagram.Interaction {
	var interactions []diagram.Interaction
	packageName := file.Name.Name

	structsMap := make(map[string]bool)
	interfacesMap := make(map[string]bool)

	// Identify structs and interfaces
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					switch s.Type.(type) {
					case *ast.StructType:
						structsMap[s.Name.Name] = true
						interactions = append(interactions, diagram.Interaction{
							Type:    "struct",
							Package: packageName,
							Name:    s.Name.Name,
						})
					case *ast.InterfaceType:
						interfacesMap[s.Name.Name] = true
						interactions = append(interactions, diagram.Interaction{
							Type:    "interface",
							Package: packageName,
							Name:    s.Name.Name,
						})
					}
				}
			}
		}
	}

	// Identify implementations of interfaces by structs via embedding
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if structType, ok := s.Type.(*ast.StructType); ok {
						for _, field := range structType.Fields.List {
							if len(field.Names) == 0 { // Embedding
								if ident, ok := field.Type.(*ast.Ident); ok {
									if interfacesMap[ident.Name] {
										interactions = append(interactions, diagram.Interaction{
											Type:    "implements",
											Package: packageName,
											From:    s.Name.Name,
											To:      ident.Name,
											Message: "implements",
										})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return interactions
}
