package main

import (
	//strings "algorithms/Strings"
	arrays  "algorithms/Arrays"
	"fmt"
)

func main() {

	//res := strings.MinimumWindowSubstring("ADOBECODEBANC", "ABC")
	//res := arrays.Subsets2([]int{1, 2, 3})
	res := arrays.CombinationSum([]int{2, 3, 4, 7}, 7)
	fmt.Println(res)
}

func DoStuff(arr *[]int) {
	*arr = (*arr)[3:6]
	// Do stuff here
}