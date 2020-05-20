package undirected

import (
	"fmt"
	"testing"
)

func TestCycleDetection(t *testing.T) {
	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=7 edge=6 has cycle",
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
		{
			"v=7 edge=5 no cycle",
			7,
			[][2]int{
				{0, 1},
				// {0, 2},
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
			fmt.Printf("%s Case, Has cycle: %v\n", testCase.name, NewCycleDetection(g).HasCycle())
		})
	}
}
