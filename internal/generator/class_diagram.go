package generator

import (
	"go-structure-diagram/pkg/diagram"
	"strings"
)

// GenerateClassDiagram creates a Mermaid class diagram string.
func GenerateClassDiagram(interactions []diagram.Interaction, packageDeps []diagram.Interaction, globalVars []diagram.Interaction) string {
	var builder strings.Builder
	builder.WriteString("classDiagram\n")

	// Add classes and interfaces
	for _, interaction := range interactions {
		switch interaction.Type {
		case "struct":
			builder.WriteString("class " + interaction.Name + "\n")
		case "interface":
			builder.WriteString("interface " + interaction.Name + "\n")
		case "implements":
			builder.WriteString(interaction.From + " --|> " + interaction.To + " : implements\n")
		}
	}

	// Add package dependencies
	for _, dep := range packageDeps {
		fromPkg := dep.From
		toPkg := dep.To
		builder.WriteString(fromPkg + " --> " + toPkg + " : imports\n")
	}

	// Add global variables
	for _, gv := range globalVars {
		builder.WriteString("class " + gv.Name + " {\n")
		builder.WriteString("    +globalVar : type\n") // Placeholder for variable type
		builder.WriteString("}\n")
	}

	return builder.String()
}
