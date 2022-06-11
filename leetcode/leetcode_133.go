package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// copy first node
	head := copyNode(node)
	// init visited map
	visited := map[int]*Node{}
	visited[node.Val] = head

	// dfs to load struct
	dfs(node, head, visited)

	return head
}

func dfs(n, curr *Node, visited map[int]*Node) {
	for _, neigh := range n.Neighbors {
		if copyNeigh, ok := visited[neigh.Val]; ok {
			curr.Neighbors = append(curr.Neighbors, copyNeigh)
			continue
		}
		cNode := copyNode(n)
		visited[neigh.Val] = cNode
		curr.Neighbors = append(curr.Neighbors, cNode)
		dfs(neigh, cNode, visited)
	}
}

func copyNode(node *Node) *Node {
	x := &Node{}
	x.Val = node.Val
	x.Neighbors = []*Node{}
	return x
}
