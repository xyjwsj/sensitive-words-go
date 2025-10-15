# Sensitive Word Filter

A Go implementation of sensitive word filtering based on DFA algorithm

## Features

- Fast detection using DFA algorithm
- Support for adding single or multiple sensitive words
- Text filtering with replacement characters
- Position tracking of sensitive words
- Case-insensitive detection
- Whitelist support
- Ignore character support
- Configurable behavior

## Installation

```bash 
go get github.com/xyjwsj/sensitive-words-go
```

## Usage

### Basic Usage

```golang
import "https://github.com/xyjwsj/sensitive-words-go"
sw := sensitive_word.NewSensitiveWord() 
sw.AddWord("敏感词")
text := "这是一段包含敏感词的内容" 
if sw.Contains(text) { 
	fmt.Println("Contains sensitive words") 
}
replaced := sw.Replace(text) 
fmt.Println(replaced) // 这是一段包含***的内容
```

### Advanced Usage

```golang
// Create with custom configuration 
config := sensitive_word.NewConfig() 
config.ReplaceChar = '#' 
config.IgnoreCase = true
sw := sensitive_word.NewSensitiveWordWithConfig(config) 
sw.AddWhiteWord("正常词汇")
if sw.Contains(text) { 
	// Handle sensitive content 
}
```

## API Reference

### Core Methods

- `NewSensitiveWord()` - Create new processor with default config
- `NewSensitiveWordWithConfig(config *Config)` - Create new processor with custom config
- `AddWord(word string)` - Add single sensitive word
- `AddWords(words []string)` - Add multiple sensitive words
- `Contains(text string) bool` - Check if text contains sensitive words
- `FindAll(text string) []string` - Find all sensitive words
- `Replace(text string) string` - Replace sensitive words with default char
- `ReplaceWithChar(text string, replaceChar rune) string` - Replace with custom char

### Configuration Methods

- `AddWhiteWord(word string)` - Add word to whitelist
- `AddWhiteWords(words []string)` - Add multiple words to whitelist
- `AddIgnoreChar(char rune)` - Add character to ignore list
- `AddIgnoreChars(chars []rune)` - Add multiple characters to ignore list
- `SetReplaceChar(char rune)` - Set replacement character
- `SetIgnoreCase(ignore bool)` - Set case sensitivity

### Advanced Methods

- `FindFirst(text string) (string, int)` - Find first sensitive word
- `FindAllWithPosition(text string) []WordPosition` - Find all with positions
- `ReplaceAndReturnCount(text string) (string, int)` - Replace and return count

## Configuration Options

- `ReplaceChar` - Character used for replacement (default: *)
- `IgnoreCase` - Whether to ignore case (default: false)
- `IgnoreChars` - Characters to ignore during detection (default: empty)
- `WhiteList` - Words that should not be treated as sensitive (default: empty)

## License

MIT
