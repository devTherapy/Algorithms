package strings

import (
	"fmt"
	"sort"
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

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	umap := make(map[rune]int)
	for _, x := range str {
		umap[x]++
	}
	for i, y := range str {
		if umap[y] == 1 {
			return i
		}
	}
	return -1
}


func isAnagram(s string, t string) bool {
	smap:= make(map[rune] int)

	for _, x := range s {
		smap[x]++
	}

	for _, y := range t {
		if smap[y] == 0 {
			return false
		}
		smap[y]--;
	}

	return true
}

func GroupAnagrams(words []string) [][]string {
	// Write your code here.
	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}
	result := make([][]string, 0)
	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result

}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func GroupAnagrams2(words []string) [][] string{
	anagramMap := make(map[string][]string)

	for _, word := range words {
		var arr = [26]int{}
		for _, c := range word{
			arr[c - 'a']++
		}
		key := fmt.Sprint(arr[:])
		anagramMap[key] = append(anagramMap[key], word)
	}

	result := make([][]string, 0)
	for _, v := range anagramMap{
		result = append(result, v)
	}
	return result
}
