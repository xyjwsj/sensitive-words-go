package sensitive_word

import (
	"strings"
)

// Replace replaces sensitive words in the text with the specified character
func (sw *SensitiveWord) Replace(text string) string {
	return sw.ReplaceWithChar(text, sw.config.ReplaceChar)
}

// ReplaceWithChar replaces sensitive words in the text with the specified character
func (sw *SensitiveWord) ReplaceWithChar(text string, replaceChar rune) string {
	runes := []rune(text)
	length := len(runes)
	result := make([]rune, length)
	copy(result, runes)

	for i := 0; i < length; i++ {
		node := sw.root
		matchLen := 0

		for j := i; j < length; j++ {
			char := runes[j]
			// Skip ignored characters if configured
			if sw.config.IgnoreChars[char] {
				continue
			}

			child, exists := node.Children[char]
			if !exists {
				break
			}

			if child.IsEnd {
				// Check whitelist
				word := string(runes[i : j+1])
				checkWord := word
				if sw.config.IgnoreCase {
					checkWord = strings.ToLower(checkWord)
				}
				if !sw.config.WhiteList[checkWord] {
					matchLen = j - i + 1
				}
			}
			node = child
		}

		if matchLen > 0 {
			for k := 0; k < matchLen; k++ {
				result[i+k] = replaceChar
			}
			i += matchLen - 1
		}
	}

	return string(result)
}

// ReplaceAndReturnCount replaces sensitive words and returns the count of replacements
func (sw *SensitiveWord) ReplaceAndReturnCount(text string) (string, int) {
	return sw.ReplaceAndReturnCountWithChar(text, sw.config.ReplaceChar)
}

// ReplaceAndReturnCountWithChar replaces sensitive words with specified char and returns count
func (sw *SensitiveWord) ReplaceAndReturnCountWithChar(text string, replaceChar rune) (string, int) {
	count := 0
	runes := []rune(text)
	length := len(runes)
	result := make([]rune, length)
	copy(result, runes)

	for i := 0; i < length; i++ {
		node := sw.root
		matchLen := 0

		for j := i; j < length; j++ {
			char := runes[j]
			// Skip ignored characters if configured
			if sw.config.IgnoreChars[char] {
				continue
			}

			child, exists := node.Children[char]
			if !exists {
				break
			}

			if child.IsEnd {
				// Check whitelist
				word := string(runes[i : j+1])
				checkWord := word
				if sw.config.IgnoreCase {
					checkWord = strings.ToLower(checkWord)
				}
				if !sw.config.WhiteList[checkWord] {
					matchLen = j - i + 1
				}
			}
			node = child
		}

		if matchLen > 0 {
			for k := 0; k < matchLen; k++ {
				result[i+k] = replaceChar
			}
			count++
			i += matchLen - 1
		}
	}

	return string(result), count
}