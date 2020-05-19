package undirected

import (
	"fmt"
	"testing"
)

func TestNewGraphMatrixFromEdge(t *testing.T) {
	relationships := []Vertex{
		{Lable: "Bob"},
		{Lable: "Alice"},
		{Lable: "Rob"},
		{Lable: "Maria"},
		{Lable: "Mark"},
	}

	cases := []struct {
		name  string
		v     int
		edges [][2]int
	}{
		{
			"v=6,edge=6",
			5,
			[][2]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewGraphMatrixFromEdge(testCase.v, testCase.edges)
			fmt.Println(g)
			for i := 0; i < len(relationships); i++ {
				it := g.Adjacency(i)
				fmt.Printf("%s: ", relationships[i].Lable.(string))
				for !it.HasNext() {
					fmt.Printf("%s ", relationships[it.Next()].Lable.(string))
				}
				fmt.Println()
			}
		})
	}

}
