package main

import (
	//strings "algorithms/Strings"
	//arrays "algorithms/Arrays"
	practice "algorithms/Practice"
	"fmt"
)

func main() {

	//res := strings.MinimumWindowSubstring("ADOBECODEBANC", "ABC")
	//res := arrays.Subsets2([]int{1, 2, 3})
	//res := arrays.CombinationSum([]int{2, 3, 4, 7}, 7)
	node := practice.TreeNode{Val: 1, Left: &practice.TreeNode{Val: 2, Left: nil, Right: nil}, Right: &practice.TreeNode{Val: 3, Left: nil, Right: nil}}
	res2 := practice.BFS(&node)
	fmt.Println(res2)
}

func DoStuff(arr *[]int) {
	*arr = (*arr)[3:6]
	// Do stuff here
}