package generator

import (
	"go-structure-diagram/pkg/diagram"
)

// Generate processes all interactions and creates the diagrams.
func Generate(interactions []diagram.Interaction) diagram.Diagrams {
	var classInteractions []diagram.Interaction
	var sequenceInteractions []diagram.Interaction
	var stateInteractions []diagram.Interaction
	var packageDependencies []diagram.Interaction
	var globalVars []diagram.Interaction

	for _, interaction := range interactions {
		switch interaction.Type {
		case "struct", "interface", "implements":
			classInteractions = append(classInteractions, interaction)
		case "func_call", "method_call":
			sequenceInteractions = append(sequenceInteractions, interaction)
		case "state_transition":
			stateInteractions = append(stateInteractions, interaction)
		case "package_dependency":
			packageDependencies = append(packageDependencies, interaction)
		case "global_var":
			globalVars = append(globalVars, interaction)
		}
	}

	return diagram.Diagrams{
		ClassDiagram:    GenerateClassDiagram(classInteractions, packageDependencies, globalVars),
		SequenceDiagram: GenerateSequenceDiagram(sequenceInteractions),
		StateDiagram:    GenerateStateDiagram(stateInteractions),
	}
}
