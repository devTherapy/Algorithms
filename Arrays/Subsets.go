package arrays


// Subsets generates all possible subsets of a given array of integers.
// It returns a slice of slices, where each inner slice is a subset of the input array.
//
// Example:
//   input := []int{1, 2, 3}
//   output := Subsets(input)
//   // output: [[], [1], [2], [1 2], [3], [1 3], [2 3], [1 2 3]]
//
// Parameters:
//   array - A slice of integers for which subsets are to be generated.
//
// Returns:
//   A slice of slices of integers, where each inner slice is a subset of the input array.
//
//dfs is a recursive function that generates all possible subsets of a given set of integers.
// It uses depth-first search (DFS) to explore all subset combinations.
//
// Parameters:
// - nums: The input slice of integers for which subsets are to be generated.
// - index: The current index in the nums slice being processed.
// - subset: The current subset being constructed.
// - result: A pointer to a slice of slices of integers where all generated subsets will be stored.
//
// The function works by including the current element in the subset and recursively calling itself
// to process the next element. It then excludes the current element from the subset and recursively
// calls itself again to process the next element. When the index reaches the length of nums, the
// current subset is added to the result.
// 
//The logic here is that we can either include or exclude an element in the subset, we can create a decision tree around this logic and explore all possible combinations of the elements in the array. Check the algorithm diagram in golang/algorithms/algorithm diagrams for a visual representation of the decision tree.
func Subsets(array []int) [][]int {
	result := [][]int{}
	subset := []int{}

	dfs(array, 0, subset, &result)
	return result
}


func dfs (nums []int, index int, subset []int, result *[][]int )  {
	if index >= len(nums) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*result = append(*result, subsetCopy)
		return
	}
	subset = append(subset, nums[index])
	dfs(nums, index + 1, subset, result)
	subset = subset[:len(subset) - 1]
	dfs(nums, index + 1, subset, result)
}



// Subsets2 generates all possible subsets of a given integer array.
// It returns a slice of slices, where each inner slice is a subset of the input array.
// The function uses an iterative approach to build subsets by adding each element
// to existing subsets.
//
// Parameters:
// - array: A slice of integers for which subsets need to be generated.
//
// Returns:
// - A slice of slices of integers, where each inner slice represents a subset of the input array.
//
// Example:
//   input: [1, 2, 3]
//   output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]
func Subsets2(array []int) [][]int{
	result := [][]int{}
	result = append(result, []int{})

	for _, v := range array{
		size := len(result)
		for j := 0; j < size; j++ {
			 subset := []int{}
			 subset = append(subset, result[j]...)
			 subset = append(subset, v)
			 result = append(result, subset)
		}
	}
	return result
}