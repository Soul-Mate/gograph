package undirected

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=7,edge=6",
			7,
			[][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 6},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphListFromEdges(testCase.v, testCase.edges)
			dfs := NewDFS(g)
			dfs.DFS()
			fmt.Printf("dfs order: %v\n", dfs.Order())
		})
	}
}
