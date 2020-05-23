package undirected

import (
	"fmt"
	"testing"
)

func TestCC(t *testing.T) {
	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=9,edge=10",
			10,
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
			cc := NewCC(g)
			fmt.Printf("Connected Component Count: %d\n", cc.ConnectedComponentCount())
			fmt.Printf("5 Connected Component: %v\n", cc.ConnectedComponent(5))
			fmt.Printf("7 Connected Component: %v\n", cc.ConnectedComponent(7))
			fmt.Printf("9 Connected Component: %v\n", cc.ConnectedComponent(9))
		})
	}
}
