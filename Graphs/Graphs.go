package graphs

type Node struct {
	Name     string
	Children []*Node
}

// DFS Graph
func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name)
	for _, c := range n.Children {
		array = c.DepthFirstSearch(array)
	}
	return array
}
