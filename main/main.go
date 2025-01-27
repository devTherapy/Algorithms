package main

import (
	strings "algorithms/Strings"
	"fmt"
)

func main() {

	//res := strings.MinimumWindowSubstring("ADOBECODEBANC", "ABC")
	res := strings.MinimumWindowSubstring2("ADOBECODEBANC", "ABC")
	fmt.Println(res)
}

func DoStuff(arr *[]int) {
	*arr = (*arr)[3:6]
	// Do stuff here
}