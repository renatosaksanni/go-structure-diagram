package packages

import (
	"go-structure-diagram/pkg/diagram"
	"go/ast"
	"strings"
)

// Analyze extracts interactions related to package dependencies.
func Analyze(file *ast.File, projectPath string) []diagram.Interaction {
	var interactions []diagram.Interaction
	packageName := file.Name.Name

	imports := make(map[string]string) // map[pkgAlias]pkgPath
	for _, imp := range file.Imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		} else {
			parts := strings.Split(importPath, "/")
			alias = parts[len(parts)-1]
		}
		imports[alias] = importPath

		// Add package_dependency interaction
		interactions = append(interactions, diagram.Interaction{
			Type:    "package_dependency",
			From:    packageName,
			To:      importPath,
			Message: "imports",
		})
	}

	return interactions
}
