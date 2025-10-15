package sensitive_word

// SensitiveWord is the main processor for sensitive word filtering
type SensitiveWord struct {
	root   *DfaNode
	config *Config
}

// NewSensitiveWord creates a new sensitive word processor with default config
func NewSensitiveWord() *SensitiveWord {
	return &SensitiveWord{
		root:   NewDfaNode(),
		config: DefaultConfig(),
	}
}

// NewSensitiveWordWithConfig creates a new sensitive word processor with custom config
func NewSensitiveWordWithConfig(config *Config) *SensitiveWord {
	return &SensitiveWord{
		root:   NewDfaNode(),
		config: config,
	}
}

// AddWord adds a sensitive word to the DFA tree
func (sw *SensitiveWord) AddWord(word string) {
	if word == "" {
		return
	}

	node := sw.root
	runes := []rune(word)

	for _, r := range runes {
		if _, exists := node.Children[r]; !exists {
			node.Children[r] = NewDfaNode()
		}
		node = node.Children[r]
	}

	node.IsEnd = true
}

// AddWords adds multiple sensitive words to the DFA tree
func (sw *SensitiveWord) AddWords(words []string) {
	for _, word := range words {
		sw.AddWord(word)
	}
}

// SetConfig sets the configuration for the sensitive word processor
func (sw *SensitiveWord) SetConfig(config *Config) {
	sw.config = config
}

// GetConfig returns the current configuration
func (sw *SensitiveWord) GetConfig() *Config {
	return sw.config
}
