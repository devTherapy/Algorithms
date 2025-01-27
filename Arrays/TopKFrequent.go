package arrays

import "sort"

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




