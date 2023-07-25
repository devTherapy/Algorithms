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
			//fill the first row in the new matrix with the first column in the old matrix etc,
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
