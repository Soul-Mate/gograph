package gograph

import (
	"bytes"
	"fmt"
)

type GraphMatrix struct {
	v int     // the graph vertex
	g [][]int // the graph
}

func NewGraphMatrixFromFile() {

}

func NewGraphMatrixFromEdge(vertices int, edges [][2]int) *GraphMatrix {
	if vertices <= 0 {
		panic("first argument must be not negative on NewGraphMatrixFromEdge()")
	}

	if len(edges) <= 0 {
		panic("second argument must be not negative on NewGraphMatrixFromEdge()")
	}

	g := new(GraphMatrix)
	g.v = vertices
	g.g = make([][]int, g.v)

	// initializtion graph
	for i := 0; i < g.v; i++ {
		g.g[i] = make([]int, g.v)
	}

	for i, edgen := 0, len(edges); i < edgen; i++ {
		v1 := edges[i][0]
		if v1 < 0 || v1 >= g.v {
			panic("vertex v invalid")
		}

		v2 := edges[i][1]
		if v2 < 0 || v2 >= g.v {
			panic("vertex v invalid")
		}

		// check self-loop edge
		if v1 == v2 {
			panic("not support self-loop edge")
		}

		//  check paralle edge
		if g.g[v1][v2] == 1 {
			panic("not support parallel edge")
		}

		g.g[v1][v2] = 1
		g.g[v2][v1] = 1
	}

	return g
}

func (g *GraphMatrix) V() int {
	return g.v
}

func (g *GraphMatrix) Adjacency(v int) GraphAdjacencyIterator {
	if v < 0 || v >= g.v {
		panic("vertex v invalid")
	}

	adjacent := make([]int, 0, g.v)
	for i := 0; i < len(g.g[v]); i++ {
		if g.g[v][i] == 1 {
			adjacent = append(adjacent, i)
		}
	}

	return NewGraphMatrixAdjacencyIterator(adjacent)
}

func (g *GraphMatrix) String() string {
	var bf bytes.Buffer
	bf.WriteString(fmt.Sprintf("Graph V: %d\n", g.v))

	// row vertex
	bf.WriteString("  ")
	for i := 0; i < g.v; i++ {
		bf.WriteString(fmt.Sprintf("%d ", i))
	}
	bf.WriteString("\n")

	for i := 0; i < g.v; i++ {
		// col vertex
		bf.WriteString(fmt.Sprintf("%d ", i))
		for j := 0; j < g.v; j++ {
			bf.WriteString(fmt.Sprintf("%d ", g.g[i][j]))
		}
		bf.WriteString("\n")
	}

	bf.WriteString(fmt.Sprintf("\nGraph Vertex Adjacent Edges:\n"))
	// adjancet edge
	for i := 0; i < g.v; i++ {
		bf.WriteString(fmt.Sprintf("Vertex %d: ", i))
		adjacencyIterator := g.Adjacency(i)
		for !adjacencyIterator.HasNext() {
			bf.WriteString(fmt.Sprintf("%d ", adjacencyIterator.Next()))
		}
		bf.WriteString("\n")
	}

	return bf.String()
}
