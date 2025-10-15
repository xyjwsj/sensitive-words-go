package sensitive_word

// Config holds the configuration for sensitive word processing
type Config struct {
	// ReplaceChar is the character used to replace sensitive words
	ReplaceChar rune

	// IgnoreCase determines whether to ignore case during detection
	IgnoreCase bool

	// IgnoreChars are characters to be ignored during detection
	IgnoreChars map[rune]bool

	// WhiteList contains words that should not be treated as sensitive
	WhiteList map[string]bool

	// SensitiveFiles
	SensitiveDirs []string
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	return &Config{
		ReplaceChar: '*',
		IgnoreCase:  false,
		IgnoreChars: make(map[rune]bool),
		WhiteList:   make(map[string]bool),
	}
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return NewConfig()
}
