package undirected

import (
	"fmt"
	"testing"
)

func TestSingleSourcePath(t *testing.T) {
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
				{0, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 6},

				{7, 8},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphListFromEdges(testCase.v, testCase.edges)
			s := NewSingleSourcePath(g, 0)
			fmt.Printf("0 -> 6: %v\n", s.Path(6))
			fmt.Printf("0 -> 5: %v\n", s.Path(5))
			fmt.Printf("0 -> 8: %v\n", s.Path(8))
		})
	}
}
