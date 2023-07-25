package main

import (
	"algorithms/arrays"
	"fmt"
)

func main() {

	array := []int{-4, -1, 3, 5, 6, 8, 11}
	ans := arrays.SortedSquaredArray(array)
	fmt.Print(ans)
}
