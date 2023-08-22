package strings

import (
	"fmt"
	"strings"
)

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	out := []rune{}
	for _, r := range str {
		shifted := r + rune(key%26)
		if shifted > 'z' {
			shifted = rune('a' - 1 + (shifted % 'z'))
		}
		out = append(out, shifted)
	}
	return string(out)
}

// Wrote a cleaner solution in C#.
func RunLengthEncoding(str string) string {
	// Write your code here.
	var builder strings.Builder
	currChar := rune(str[0])
	count := 0
	for _, r := range str {
		if r == currChar && count < 9 {
			count++
		} else {
			builder.WriteString(fmt.Sprintf("%d%s", count, string(currChar)))
			currChar = r
			count = 1
		}

	}
	if count > 0 {
		builder.WriteString(fmt.Sprintf("%d%s", count, string(currChar)))
	}
	return builder.String()
}

func CommonCharacters(strings []string) []string {
	// Write your code here.
	counts := make(map[rune]int)
	for _, str := range strings {
		seen := make(map[rune]bool)
		for _, r := range str {
			if !seen[r] {
				seen[r] = true
				counts[r]++
			}
		}
	}
	result := make([]string, 0)

	for char, count := range counts {
		if count == len(strings) {
			result = append(result, string(char))
		}
	}
	return result
}

// not in c#
func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	umap := make(map[rune]int)

	for _, x := range characters {
		umap[x]++
	}
	for _, y := range document {
		if umap[y] == 0 {
			return false
		}
		umap[y]--
	}

	return true
}
