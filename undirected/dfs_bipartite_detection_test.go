package undirected

import (
	"testing"
)

func TestBipartiteDetection(t *testing.T) {
	cases := []struct {
		name        string
		v           int
		edges       [][2]int
		isBipartite bool
	}{
		{
			"v-7,e-6,is-bipartite-graph",
			7,
			[][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 6},
			},
			true,
		},
		{
			"v-7,e-7,not-bipartite-graph",
			7,
			[][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 6},
				{4, 6},
			},
			false,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphListFromEdges(testCase.v, testCase.edges)
			bi := NewBipartiteDetection(g)

			if bi.IsBipartite() != testCase.isBipartite {
				t.Errorf("%s Case faild, The graph bipartite is: %v, got: %v\n", testCase.name,
					testCase.isBipartite, bi.IsBipartite())
			}
		})
	}
}
