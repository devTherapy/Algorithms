package arrays

// CombinationSum returns all possible combinations of elements that sum up to the target value.
// It uses a depth-first search (DFS) approach to explore all possible combinations.
//
// Parameters:
// - array: A slice of integers from which combinations are to be generated.
// - target: The target value that the combinations should sum up to.
//
// Returns:
// A slice of slices of integers, where each inner slice is a combination of elements that sum up to the target value.
//
// Example:
//   array := []int{2, 3, 6, 7}
//   target := 7
//   result = [[2 2 3] [7]]
//   
//
// Just like in the subsets example, we use a depth-first search (DFS) approach to explore all possible combinations of elements that sum up to the target value. We start by adding the current element to the subset and recursively calling the function with the updated total. If the total equals the target value, we add the subset to the result. If the total exceeds the target value or we reach the end of the array, we backtrack and remove the last element from the subset.
//
// The logic here is that we can either include or exclude an element in the subset, and we can create a decision tree around this logic to explore all possible combinations of elements in the array. Check the algorithm diagram in the golang/algorithms/algorithm diagrams folder for a visual representation of the decision tree. However, the questiom specifies we are allowed to reuse elements, so we can include the same element multiple times in the combination. This is why in the left branch of the decision tree, we call the function with the same index, but in the right branch, we call the function with the next index.
func CombinationSum(array []int, target int) [][]int{
	result := [][]int{}
	subset := []int{}
	var dfs func(index int, total int)
	dfs = func(index int, total int) {
		if total == target {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			result = append(result, subsetCopy)
			return
		}

		if total > target || index >= len(array) {
			return
		}

		subset = append(subset, array[index])
		dfs(index, total + array[index])
		subset = subset[:len(subset) - 1]
		dfs(index + 1, total)
	}
	dfs(0, 0)
	return result
}