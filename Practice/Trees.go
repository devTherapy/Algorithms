package practice

import "container/list"


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// PreOrderTraversal performs a pre-order traversal of the given binary tree.
// It returns a slice of integers representing the order of traversal.
//
// Parameters:
// - root: A pointer to the root node of the binary tree.
//
// Returns:
// A slice of integers representing the order of traversal.
//
// Example:
// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
// result := PreOrderTraversal(root) // result will be [1, 2, 3]
//
// it is called pre-order because the root node is visited before the left and right subtrees.
func PreOrderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        result = append(result, node.Val)
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return result
}

// InOrderTraversal performs an in-order traversal of the given binary tree.
// It returns a slice of integers representing the order of traversal.
//
// Parameters:
// - root: A pointer to the root node of the binary tree.
//
// Returns:
// A slice of integers representing the order of traversal.
//
// Example:
// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
// result := InOrderTraversal(root) // result will be [2, 1, 3]
//
// it is called in-order because the root node is visited in between the left and right subtrees.
func InOrderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        result = append(result, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return result
}

// PostOrderTraversal performs a post-order traversal of the given binary tree.
// It returns a slice of integers representing the order of traversal.
//
// Parameters:
// - root: A pointer to the root node of the binary tree.
//
// Returns:
// A slice of integers representing the order of traversal.
//
// Example:
// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
// result := PostOrderTraversal(root) // result will be [2, 3, 1]
//
// it is called post-order because the root node is visited after the left and right subtrees.
func PostOrderTraversal(root *TreeNode) []int {
    result := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        dfs(node.Right)
        result = append(result, node.Val)
    }
    dfs(root)
    return result
}

// BFS performs a breadth-first search traversal of the given binary tree.
// It returns a slice of integers representing the order of traversal.
//
// Parameters:
// - root: A pointer to the root node of the binary tree.
//
// Returns:
// A slice of integers representing the order of traversal.
//
// Example:
// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
// result := BFS(root) // result will be [1, 2, 3]
//
// In a breadth-first search traversal, the nodes are visited level by level, starting from the root node.
func BFS(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    result := []int{}
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        element := queue.Front()
        node := element.Value.(*TreeNode)
        queue.Remove(element)

        result = append(result, node.Val)

        if node.Left != nil {
            queue.PushBack(node.Left)
        }
        if node.Right != nil {
            queue.PushBack(node.Right)
        }
    }

    return result
}


