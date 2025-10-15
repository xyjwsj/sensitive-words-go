package sensitive_word

// DfaNode represents a node in the DFA tree
type DfaNode struct {
	Children map[rune]*DfaNode
	IsEnd    bool
}

// NewDfaNode creates a new DFA node
func NewDfaNode() *DfaNode {
	return &DfaNode{
		Children: make(map[rune]*DfaNode),
		IsEnd:    false,
	}
}
