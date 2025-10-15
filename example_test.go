package sensitive_word

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	sw := NewSensitiveWord()
	sw.AddWord("敏感词")
	text := "这是一段包含敏感词的内容"
	if sw.Contains(text) {
		fmt.Println("Contains sensitive words")
	}
	replaced := sw.Replace(text)
	fmt.Println(replaced) // 这是一段包含***的内容
}

func TestTextAdv(t *testing.T) {
	text := "这是一段包含敏感词的内容"
	config := NewConfig()
	config.ReplaceChar = '#'
	config.IgnoreCase = true
	sw := NewSensitiveWordWithConfig(config)
	sw.AddWhiteWord("正常词汇")
	if sw.Contains(text) {
		// Handle sensitive content
	}
}
