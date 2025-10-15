package sensitive_word

import (
	"strings"
)

// Contains checks if the text contains any sensitive words
func (sw *SensitiveWord) Contains(text string) bool {
	if sw.config.IgnoreCase {
		text = strings.ToLower(text)
	}

	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length; i++ {
		// Skip ignored characters if configured
		if sw.config.IgnoreChars[runes[i]] {
			continue
		}

		node := sw.root
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
				if sw.config.IgnoreCase {
					word = strings.ToLower(word)
				}
				if !sw.config.WhiteList[word] {
					return true
				}
			}
			node = child
		}
	}

	return false
}

// FindAll finds all sensitive words in the text
func (sw *SensitiveWord) FindAll(text string) []string {
	var result []string

	if sw.config.IgnoreCase {
		text = strings.ToLower(text)
	}

	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length; i++ {
		// Skip ignored characters if configured
		if sw.config.IgnoreChars[runes[i]] {
			continue
		}

		node := sw.root
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
				word := string(runes[i : j+1])
				// Check whitelist
				checkWord := word
				if sw.config.IgnoreCase {
					checkWord = strings.ToLower(checkWord)
				}
				if !sw.config.WhiteList[checkWord] {
					result = append(result, word)
				}
			}
			node = child
		}
	}

	return result
}

// FindFirst finds the first sensitive word in the text and its position
func (sw *SensitiveWord) FindFirst(text string) (string, int) {
	if sw.config.IgnoreCase {
		text = strings.ToLower(text)
	}

	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length; i++ {
		// Skip ignored characters if configured
		if sw.config.IgnoreChars[runes[i]] {
			continue
		}

		node := sw.root
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
				word := string(runes[i : j+1])
				// Check whitelist
				checkWord := word
				if sw.config.IgnoreCase {
					checkWord = strings.ToLower(checkWord)
				}
				if !sw.config.WhiteList[checkWord] {
					return word, i
				}
			}
			node = child
		}
	}

	return "", -1
}

// WordPosition represents the position information of a sensitive word
type WordPosition struct {
	Word     string
	StartPos int
	EndPos   int
}

// FindAllWithPosition finds all sensitive words with their positions
func (sw *SensitiveWord) FindAllWithPosition(text string) []WordPosition {
	var result []WordPosition

	if sw.config.IgnoreCase {
		text = strings.ToLower(text)
	}

	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length; i++ {
		// Skip ignored characters if configured
		if sw.config.IgnoreChars[runes[i]] {
			continue
		}

		node := sw.root
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
				word := string(runes[i : j+1])
				// Check whitelist
				checkWord := word
				if sw.config.IgnoreCase {
					checkWord = strings.ToLower(checkWord)
				}
				if !sw.config.WhiteList[checkWord] {
					result = append(result, WordPosition{
						Word:     word,
						StartPos: i,
						EndPos:   j,
					})
				}
			}
			node = child
		}
	}

	return result
}