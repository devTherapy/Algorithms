package main

import (
	//strings "algorithms/Strings"
	arrays  "algorithms/Arrays"
	"fmt"
)

func main() {

	//res := strings.MinimumWindowSubstring("ADOBECODEBANC", "ABC")
	res := arrays.MaximumSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(res)
}

func DoStuff(arr *[]int) {
	*arr = (*arr)[3:6]
	// Do stuff here
}