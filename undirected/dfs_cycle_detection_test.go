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
			c := NewCycleDetection(g)
			fmt.Printf("start Vertex 0, has cycle: %v\n", c.Detection(0))
			fmt.Printf("start Vertex 1, has cycle: %v\n", c.Detection(1))
			fmt.Printf("start Vertex 3, has cycle: %v\n", c.Detection(3))
			fmt.Printf("start Vertex 2, has cycle: %v\n", c.Detection(2))
			fmt.Printf("start Vertex 4, has cycle: %v\n", c.Detection(4))
			fmt.Printf("start Vertex 6, has cycle: %v\n", c.Detection(6))
			fmt.Printf("start Vertex 5, has cycle: %v\n", c.Detection(5))
		})
	}
}
