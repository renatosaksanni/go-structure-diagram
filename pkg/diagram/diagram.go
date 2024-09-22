package diagram

// Interaction represents a relationship or interaction between code elements.
type Interaction struct {
	Type      string // e.g., "struct", "interface", "func", "method_call", "package_dependency", "global_var", "implements"
	Package   string // Package where the element resides
	Name      string // Name of the element
	From      string // Source of the interaction
	To        string // Target of the interaction
	Message   string // Description of the interaction (used in sequence diagrams)
	StateFrom string // Used for state diagrams
	StateTo   string // Used for state diagrams
	Event     string // Used for state diagrams
}
