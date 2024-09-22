package generator

import (
	"go-structure-diagram/pkg/diagram"
	"strings"
)

// GenerateStateDiagram creates a Mermaid state diagram string.
// This is a basic implementation and may need adjustments based on actual state transitions.
func GenerateStateDiagram(interactions []diagram.Interaction) string {
	var builder strings.Builder
	builder.WriteString("stateDiagram-v2\n")

	for _, interaction := range interactions {
		if interaction.Type == "state_transition" {
			builder.WriteString(interaction.StateFrom + " --> " + interaction.StateTo + ": " + interaction.Event + "\n")
		}
	}

	return builder.String()
}
