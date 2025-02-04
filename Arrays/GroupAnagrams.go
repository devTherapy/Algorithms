package arrays

// GroupAnagrams groups an array of strings into anagrams.
//
// An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Parameters:
//   arr []string - An array of strings to be grouped into anagrams.
//
// Returns:
//   [][]string - A 2D array where each sub-array contains strings that are anagrams of each other.
//
//TC O(N), SC O(N)
func GroupAnagrams(arr []string) [][]string {
	// Create a map whose key is an array of 26 integers representing a word  and value is an array of strings representing the anagrams of that word.
	umap := make(map[[26]int][]string)

	for _, word := range arr {
		// Create an array of 26 integers to represent the word.
		wordArr := [26]int{}
		for _, char := range word {
			//Get the position of the character in relation to 'a' and add one to get the position in the alphabet.
			wordArr[char-'a']++
		}
		// Append the word to the list of anagrams for that word using the array as the key.
		umap[wordArr] = append(umap[wordArr], word)
	}

	res := make([][]string, 0, len(umap))
	for _, v := range umap {
		res = append(res, v)
	}
	return res
}
