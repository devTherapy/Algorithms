
package arrays

import "sort"

// TopKFrequentElements returns the k most frequent elements from the given array.
// It uses a map to count the frequency of each element and then sorts the elements
// based on their frequency in descending order.
//
// Parameters:
// - array: A slice of integers from which the k most frequent elements are to be found.
// - k: An integer representing the number of top frequent elements to return.
//
// Returns:
// A slice of integers containing the k most frequent elements from the input array.
//
// Example:
// array := []int{1, 1, 1, 2, 2, 3}
// k := 2
// result := TopKFrequentElements(array, k) // result will be [1, 2]
func TopKFrequentElements(array []int, k int) []int {
	umap := make(map[int]int)
	for _, v := range array {
		umap[v]++
	}
	arrayofFreq := make([][2]int, 0, len(array))
	for k, v := range umap {
		arrayofFreq = append(arrayofFreq, [2]int{k, v})

		sort.Slice(arrayofFreq, func(i, j int) bool {
			return arrayofFreq[i][1] > arrayofFreq[j][1]
		})
	}

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res = append(res, arrayofFreq[i][0])
	}

	return res
}

// TopKFrequent2 returns the k most frequent elements from the given array.
// It uses a map to count the frequency of each element and then organizes the elements
// into a frequency array where the index represents the frequency.
//
// Parameters:
// - array: A slice of integers from which the k most frequent elements are to be found.
// - k: An integer representing the number of top frequent elements to return.
//
// Returns:
// A slice of integers containing the k most frequent elements from the input array.
//
// Example:
// array := []int{1, 1, 1, 2, 2, 3}
// k := 2
// result := TopKFrequent2(array, k) // result will be [1, 2]
func TopKFrequent2(array []int, k int) []int{
	frequencyMap := make(map[int]int)
	frequencyArray := make([][]int, 0, len(array) + 1)

	for _, v := range array{
		frequencyMap[v]++
	}
	
	for k, v := range frequencyMap{
		frequencyArray[v] = append(frequencyArray[v], k)
	}

	res:= []int{}
	for i := len(frequencyArray) - 1; i > 0; i--{
		res = append(res, frequencyArray[i]...)

		if len(res) == k {
			return res
		}
	}
	return res
}




