package strings

import "math"

// MinimumWindowSubstring finds the minimum window in a string s which will contain all the characters in string t.
//
// Parameters:
//   s string - The string to search for the minimum window.
//   t string - The string to search for in the minimum window.
//
// Returns:
//   string - The minimum window in string s which will contain all the characters in string t.
//
//TC O(N^2), SC O(N)
//
//This function works by iterating through the string s and checking if the substring contains all the characters in t. If the substring contains all the characters in t and the length of the substring is less than the current minimum substring length then update the minimum substring length.
//
// This is a Brute Force solution and can be optimized using a sliding window approach.
func MinimumWindowSubstring(s string, t string) string {
 
	//store the count of the substring t
   tCount := make(map[rune]int)

   for  _, c := range t {
	tCount[c]++
   }
   //initialize a default substring left and right pointers
   resIndexes := [2]int{-1, -1}

   //initialize a default substring length
   resLength := math.MaxInt

  //iterate through the string s and check if the substrings contains all the characters in t
   for i := range s {
	sCount:= make(map[rune]int)
	for j := i; j < len(s); j++ {
		sCount[rune(s[j])]++
		containsAll := true
		//check if the substring contains all the characters in t
		for k, v := range tCount{
			if sCount[k] < v {
				containsAll = false
				break
			}
		}
		//if the substring contains all the characters in t and the length of the substring is less than the current minimum substring length then update the minimum substring length
		if containsAll && (j - i + 1) < resLength {
			resIndexes[0] = i
			resIndexes[1] = j
			resLength = j - i + 1
		}
	}
	
   }
   if resLength == -1 {
	return ""
   }
   return s[resIndexes[0]:resIndexes[1] + 1]
}

// MinimumWindowSubstring2 finds the minimum window in a string s which will contain all the characters in string t.
//
// Parameters:
//   s string - The string to search for the minimum window.
//   t string - The string to search for in the minimum window.
//
// Returns:
//   string - The minimum window in string s which will contain all the characters in string t.
//
//TC O(N), SC O(N)
//
//This function works by using a sliding window approach to find the minimum window in a string s which will contain all the characters in string t.
//
// The function initializes a map to store the count of the substring t and a map to store the count of the current window. It then iterates through the string s and updates the window count. If the window count is equal to the count of the substring t, it updates the left pointer of the window and checks if the current substring is the minimum substring. If the current substring is the minimum substring, it updates the minimum substring indices.
func MinimumWindowSubstring2(s string, t string) string {
	//store the count of the substring t
	tCount := make(map[rune]int)
	window := make(map[rune]int)

	for  _, c := range t {
		tCount[c]++
	}
	//initialize a default substring indices
	resIndexes := [2]int{-1, -1}

	//initialize the left and right pointers
	resLength := math.MaxInt 
	l := 0
	//HAVE is the number of characters in the window that are equal to the characters in the substring t while NEED is the number of characters in the substring t
	have, need := 0, len(tCount)
	for i, v := range s {
		window[v]++
		if window[v] == tCount[v] {
			have++
		}

		for need == have {
			subStringLength := i - l + 1
			if (subStringLength < resLength)  {
				resLength = subStringLength
				resIndexes[0] = l
				resIndexes[1] = i
			}
			leftChar := s[l]
			//while we have what we need, we keep removing the leftmost characters from the window until we are sure we have a smallest substring that contains all the characters in the substring t 
			window[rune(leftChar)]--
			if window[rune(leftChar)] < tCount[rune(leftChar)] {
				have--
			}
			l++
		}
	}
	return s[resIndexes[0]: resIndexes[1] + 1]
}