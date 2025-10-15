package sensitive_word

import "strings"

// AddWhiteWord adds a word to the whitelist
func (sw *SensitiveWord) AddWhiteWord(word string) {
	if sw.config.IgnoreCase {
		word = strings.ToLower(word)
	}
	sw.config.WhiteList[word] = true
}

// AddWhiteWords adds multiple words to the whitelist
func (sw *SensitiveWord) AddWhiteWords(words []string) {
	for _, word := range words {
		sw.AddWhiteWord(word)
	}
}

// RemoveWhiteWord removes a word from the whitelist
func (sw *SensitiveWord) RemoveWhiteWord(word string) {
	if sw.config.IgnoreCase {
		word = strings.ToLower(word)
	}
	delete(sw.config.WhiteList, word)
}

// AddIgnoreChar adds a character to be ignored during detection
func (sw *SensitiveWord) AddIgnoreChar(char rune) {
	sw.config.IgnoreChars[char] = true
}

// AddIgnoreChars adds multiple characters to be ignored during detection
func (sw *SensitiveWord) AddIgnoreChars(chars []rune) {
	for _, char := range chars {
		sw.config.IgnoreChars[char] = true
	}
}

// RemoveIgnoreChar removes a character from the ignore list
func (sw *SensitiveWord) RemoveIgnoreChar(char rune) {
	delete(sw.config.IgnoreChars, char)
}

// SetReplaceChar sets the character used for replacement
func (sw *SensitiveWord) SetReplaceChar(char rune) {
	sw.config.ReplaceChar = char
}

// SetIgnoreCase sets whether to ignore case during detection
func (sw *SensitiveWord) SetIgnoreCase(ignore bool) {
	sw.config.IgnoreCase = ignore
}