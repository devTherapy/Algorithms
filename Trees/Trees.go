package trees

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	smallestDifference := math.MaxFloat64
	closestValue := 0
	currentNode := tree
	for tree != nil {
		currentDifference := math.Abs(float64(target - currentNode.Value))
		if currentDifference < smallestDifference {
			smallestDifference = currentDifference
			closestValue = currentNode.Value
		}
		if target > tree.Value {
			tree = tree.Right
		} else if target < tree.Value {
			tree = tree.Left
		} else {
			break
		}
	}
	return closestValue
}

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	var sums []int
	CalculateBranchSums(root, 0, &sums)
	return sums
}

func CalculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {

	if node == nil {
		return
	}
	newRunningSum := runningSum + node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, newRunningSum)
	}

	CalculateBranchSums(node.Left, newRunningSum, sums)
	CalculateBranchSums(node.Right, newRunningSum, sums)

}

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	sumOfNodeDepths := 0
	currentDepth := 0

	SumNodeDepths(root, currentDepth, &sumOfNodeDepths)
	return sumOfNodeDepths
}

func SumNodeDepths(node *BinaryTree, currentDepth int, sumOfNodeDepths *int) {
	if node == nil {
		return
	}
	*sumOfNodeDepths += currentDepth
	SumNodeDepths(node.Left, currentDepth+1, sumOfNodeDepths)
	SumNodeDepths(node.Right, currentDepth+1, sumOfNodeDepths)
}
