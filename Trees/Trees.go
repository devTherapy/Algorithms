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

	closestValue := tree.Value;

	for tree != nil {
		if math.Abs(float64(closestValue) - float64(target)) > math.Abs((float64(tree.Value) - float64(target) )){
			closestValue = tree.Value;
		}

		if target > tree.Value {
			tree = tree.Right
		}	else if target < tree.Value {
			tree = tree.Left
		} else{
			break;
		}
	}
	return closestValue;
}
