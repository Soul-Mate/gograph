package gograph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=9,edge=10",
			9,
			[][2]int{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 6},
				{2, 3},
				{2, 5},
				{3, 4},
				{5, 4},
				{5, 6},
				{7, 8},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphListFromEdges(testCase.v, testCase.edges)
			dfs := NewDFS(g)
			dfs.DFS()
		})
	}
}
