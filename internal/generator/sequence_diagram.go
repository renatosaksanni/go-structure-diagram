package generator

import (
	"go-structure-diagram/pkg/diagram"
	"strings"
)

// GenerateSequenceDiagram creates a Mermaid sequence diagram string.
func GenerateSequenceDiagram(interactions []diagram.Interaction) string {
	var builder strings.Builder
	builder.WriteString("sequenceDiagram\n")

	for _, interaction := range interactions {
		switch interaction.Type {
		case "func_call", "method_call":
			builder.WriteString(interaction.From + "->>" + interaction.To + ": " + interaction.Message + "\n")
		}
	}

	return builder.String()
}
