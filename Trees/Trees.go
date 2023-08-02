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

// recursion solution
func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	sumOfNodeDepths := 0
	currentNodeDepth := 0

	SumNodeDepths(root, currentNodeDepth, &sumOfNodeDepths)
	return sumOfNodeDepths
}

func SumNodeDepths(node *BinaryTree, currentNodeDepth int, sumOfNodeDepths *int) {
	if node == nil {
		return
	}
	*sumOfNodeDepths += currentNodeDepth
	SumNodeDepths(node.Left, currentNodeDepth+1, sumOfNodeDepths)
	SumNodeDepths(node.Right, currentNodeDepth+1, sumOfNodeDepths)
}

// iterative solution
func NodeDepths_Iterative(root *BinaryTree) int {
	sumOfNodeDepths := 0

	stack := []struct {
		node  *BinaryTree
		depth int
	}{}

	stack = append(stack, struct {
		node  *BinaryTree
		depth int
	}{root, 0})

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := item.node
		depth := item.depth
		sumOfNodeDepths += depth
		if node.Left != nil {
			stack = append(stack, struct {
				node  *BinaryTree
				depth int
			}{node.Left, depth + 1})
		}

		if node.Right != nil {
			stack = append(stack, struct {
				node  *BinaryTree
				depth int
			}{node.Right, depth + 1})
		}
	}
	return sumOfNodeDepths
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	if tree.Value > 0 {
		return tree.Value
	}

	leftValue := EvaluateExpressionTree(tree.Left)
	rightValue := EvaluateExpressionTree(tree.Right)

	switch tree.Value {
	case -1:
		return leftValue + rightValue
	case -2:
		return leftValue - rightValue
	case -3:
		return leftValue / rightValue
	default:
		return leftValue * rightValue
	}
}
