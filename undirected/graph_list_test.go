package undirected

import (
	"fmt"
	"testing"
)

func TestNewGraphListFromEdges(t *testing.T) {
	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=7,edge=9",
			7,
			[][2]int{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 6},
				{3, 4},
				{2, 3},
				{5, 2},
				{5, 4},
				{6, 5},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphListFromEdges(testCase.v, testCase.edges)
			fmt.Println(g)
			fmt.Println(g.IsConnected(0, 1), g.IsConnected(0, 5))
		})
	}
}
