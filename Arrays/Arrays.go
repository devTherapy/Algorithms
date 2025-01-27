package arrays

import (
	"math"
	"sort"
)

// 0(N) tc, 0(N) sc
func TwoSum(array []int, target int) []int {

	storedValues := map[int]struct{}{}

	for _, currentNum := range array {
		diff := target - currentNum

		if _, ok := storedValues[diff]; ok {
			return []int{currentNum, diff}
		}

		storedValues[currentNum] = struct{}{}
	}
	return []int{}
}

// func TwoSum2(array []int, target int) []int {
// 	sort.Ints(array)
// 	left := 0
// 	right := len(array) - 1
// 	expectedSum := array[left] + array[right];
// 	for left < right {

// 		if expectedSum == target {
//        		 return []int{array[left], array[right]}
// 		}
// 		if expectedSum < target {
// 			left++;
// 		}

// 		if expectedSum > target {
// 			right++;
// 		}

// 	}
// 	return[]int{}
// }

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	outputArray := make([]int, len(array))
	firstIndex := 0
	lastIndex := len(array) - 1
	counter := len(array) - 1

	for i := 0; i < len(array); i++ {
		if math.Abs(float64(array[firstIndex])) > math.Abs(float64(array[lastIndex])) {
			outputArray[counter-i] = array[firstIndex] * array[firstIndex]
			firstIndex++
		} else {
			outputArray[counter-i] = array[lastIndex] * array[lastIndex]
			lastIndex--
		}
	}
	return outputArray
}

// func TournamentWinner(competitions [][]string, results []int) string {
// 	// Write your code here.
// 	competitionWinner := ""
// 	winnerDict := map[string]int{}
// 	maxScore := 0
// 	for i := 0; i < len(competitions); i++ {
// 		competition := competitions[i]
// 		var pos int32
// 		score := 0
// 		if results[i] == 0 {
// 			pos = 1
// 		} else {
// 			pos = 0
// 		}
// 		roundWinner := competition[pos]

// 		if _, found := winnerDict[roundWinner]; found {
// 			winnerDict[roundWinner] = winnerDict[roundWinner] + 3
// 			score = winnerDict[roundWinner]
// 		} else {
// 			winnerDict[roundWinner] = 3
// 			score = winnerDict[roundWinner]
// 		}

// 		if score > maxScore {
// 			maxScore = score
// 			competitionWinner = roundWinner
// 		}
// 	}
// 	return competitionWinner
// }

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	winner := ""
	scores := map[string]int{}
	for i := range competitions {
		roundWinner := competitions[i][1-results[i]]
		scores[roundWinner] += 3
		if scores[roundWinner] > scores[winner] {
			winner = roundWinner
		}
	}
	return winner
}

func TransposeMatrix(matrix [][]int) [][]int {
	//create a new array with the length of the matrix rows since they will be the new columns
	newMatrix := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		newMatrix[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			//fill the first row in the new array with the first column in the old matrix etc,
			newMatrix[i][j] = matrix[j][i]
		}
	}

	return newMatrix
}

func NonConstructibleChange(coins []int) int {
	// Write your code here.

	sort.Ints(coins)
	change := 0

	for i := range coins {
		if coins[i] > change+1 {
			return change + 1
		}
		change += coins[i]
	}
	return change + 1
}

// "queries": [3, 2, 1, 2, 6]
func MinimumWaitingTimes(queries []int) int {

	sort.Ints(queries)
	totalWaitingTime := 0

	for i, duration := range queries {
		queriesLeft := len(queries) - (i + 1)
		totalWaitingTime += queriesLeft * duration
	}
	return totalWaitingTime
}

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	// Write your code here.

	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	redIsBackRow := redShirtHeights[len(redShirtHeights)-1] > blueShirtHeights[len(blueShirtHeights)-1]
	for i := range redShirtHeights {
		if redIsBackRow && redShirtHeights[i] <= blueShirtHeights[i] {
			return false
		} else if !redIsBackRow && blueShirtHeights[i] <= redShirtHeights[i] {
			return false
		}
	}
	return true
}

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	// Write your code here.
	sort.Ints(blueShirtSpeeds)
	sort.Ints(redShirtSpeeds)
	totalSpeed := 0

	for i := range redShirtSpeeds {
		j := i
		if fastest {
			// point to the last speed in the redshirt speeds
			j = len(redShirtSpeeds) - 1 - i
		}
		totalSpeed += int(math.Max(float64(blueShirtSpeeds[i]), float64(redShirtSpeeds[j])))
	}
	return totalSpeed
}

func OptimalFreelancing(jobs []map[string]int) int {
	// Write your code here.
	profitsTaken := make([]int, 7)

	for i := range jobs {
		deadline := jobs[i]["deadline"]
		if deadline > 7 {
			deadline = 7
		}
		currentPayment := jobs[i]["payment"]

		for j := deadline - 1; j >= 0; j-- {
			if profitsTaken[j] < currentPayment && currentPayment != 0 {
				temp := profitsTaken[j]
				profitsTaken[j] = currentPayment
				currentPayment = temp
			}
		}
	}

	sum := 0
	for i := range profitsTaken {
		sum += i
	}
	return sum
}

type SpecialArray []interface{}

func ProductSum(array []interface{}) int {
	// Write your code here.

	return ProductSums(array, 1)
}

func ProductSums(array []interface{}, depth int) int {
	sum := 0
	for _, item := range array {
		t, ok := item.(int)
		if ok {
			sum += t
		} else {
			sum += ProductSums(item.(SpecialArray), depth+1)
		}
	}
	return sum * depth
}

func BinarySearch(array []int, target int) int {
	// Write your code here.
	left, right := 0, len(array)-1

	for left <= right {
		midPoint := (left + right) / 2

		if array[midPoint] == target {
			return midPoint
		} else if array[midPoint] < target {
			left = midPoint + 1
		} else {
			right = midPoint - 1
		}
	}
	return -1
}

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	min, mid, max := math.MinInt32, math.MinInt32, math.MinInt32

	for _, num := range array {
		if num > max {
			min = mid
			mid = max
			max = num
		} else if num > mid {
			min = mid
			mid = num
		} else if num > min {
			min = num
		}
	}
	return []int{min, mid, max}
}

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	list := [][]int{}
	sort.Ints(array)
	for i, val := range array {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := val + array[left] + array[right]
			if currentSum == target {
				list = append(list, []int{val, array[left], array[right]})
				left++
				right--
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return list
}

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	left, right := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for left < len(array1) && right < len(array2) {
		first, second := array1[left], array2[right]
		if first < second {
			current = second - first
			left++
		} else if second < first {
			current = first - second
			right++
		} else {
			return []int{first, second}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	left, right := 0, len(array)-1
	for left < right {
		for left < right && array[right] == toMove {
			right--
		}
		if array[left] == toMove {
			array[left], array[right] = array[right], array[left]
		}
		left++
	}
	return array
}

func IsMonotonic(array []int) bool {
	// Write your code here.
	isNonDecreasing := true
	isNonIncreasing := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		if array[i] > array[i-1] {
			isNonIncreasing = false
		}
	}
	return isNonDecreasing || isNonIncreasing
}

func LongestPeak(array []int) int {
	// Write your code here.

	if len(array) <= 2 {
		return 0
	}

	longestPeakLength := 0
	i := 1
	for i < len(array) {
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			continue
		}

		leftIndex := i - 2
		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			leftIndex--
		}

		rightIndex := i + 2

		for rightIndex <= len(array) && array[rightIndex] < array[rightIndex-1] {
			rightIndex++
		}
		currentPeakLength := rightIndex - leftIndex - 1
		longestPeakLength = int(math.Max(float64(longestPeakLength), float64(currentPeakLength)))
		i = rightIndex
	}
	return longestPeakLength
}

// Solution without using division  [7, 1,, 2, 3, 4, 5, 6]
func ArrayOfProducts(array []int) []int {
	// Write your code here.
	product := 1
	//compute the left product
	result := make([]int, len(array))
	for i, val := range array {
		result[i] = product
		product = product * val
	}
	product = 1
	//compute the ridht product
	for i := len(array) - 1; i >= 0; i-- {
		result[i] = result[i] * product
		product = product * result[i]
	}
	return result
}

func ProductExceptSelf(array []int) []int {
	// Write your code here.
	leftProduct := make([]int, len(array))
	rightProduct := make([]int, len(array))
	result := make([]int, len(array))
	//compute the left product
	product := 1
	for i, val := range array{
		leftProduct[i] = product
		product = product * val
	}
	//compute the right product
	product = 1
	for i := len(array) - 1; i >= 0; i--{
		rightProduct[i] = product
		product = product * array[i]
	}
	//compute the product except self
	for i := range array{
		result[i] = leftProduct[i] * rightProduct[i]
	}
	return result
}

func ProductExceptSelf2(array []int) []int {
	// Write your code here.
	zeroCount := 0
	product := 1
	for _, val := range array{
		if val == 0{
			zeroCount++
		}else{
			product = product * val
		}
	}
	if zeroCount > 1{
		return make([]int, len(array))		
	}

	result := make([]int, len(array))
	for i, val := range array{
		if zeroCount == 1{
			if val == 0{
				result[i] = product
			}else{
				result[i] = 0
			}
		}else{
			result[i] = product
	}
}
	return result
}
