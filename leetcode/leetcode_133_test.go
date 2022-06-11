package leetcode

import (
	"fmt"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	var ns []*Node
	for i := 0; i < 4; i++ {
		ns = append(ns, &Node{
			Val:       i + 1,
			Neighbors: nil,
		})
	}
	for i, n := range ns {
		n.Neighbors = append(n.Neighbors, ns[(i+1+4)%4])
		n.Neighbors = append(n.Neighbors, ns[(i-1+4)%4])
	}
	fmt.Println(cloneGraph(ns[0]))
}
