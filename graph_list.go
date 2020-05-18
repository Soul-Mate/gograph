package gograph

import (
	"bytes"
	"container/list"
	"fmt"
)

type GraphList struct {
	v int
	g []*list.List
}

func NewGraphListFromEdges(vertices int, edges [][2]int) *GraphList {
	if vertices <= 0 {
		panic("first argument must be not negative on NewGraphMatrixFromEdge()")
	}

	if len(edges) <= 0 {
		panic("second argument must be not negative on NewGraphMatrixFromEdge()")
	}

	g := new(GraphList)
	g.v = vertices
	g.g = make([]*list.List, g.v)

	for i := 0; i < g.v; i++ {
		g.g[i] = list.New()
	}

	for i, edgen := 0, len(edges); i < edgen; i++ {
		from := edges[i][0]
		if from < 0 || from >= g.v {
			panic("vertex v invalid")
		}

		to := edges[i][1]
		if to < 0 || to >= g.v {
			panic("vertex v invalid")
		}

		if from == to {
			panic("not support self-loop edge")
		}

		l := g.g[from]
		parallelEdgeFlag := false
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == to {
				parallelEdgeFlag = true
				break
			}
		}

		if parallelEdgeFlag {
			panic("not support parallel edge")
		}

		// add to list
		l.PushBack(to)
		g.g[to].PushBack(from)
	}

	return g
}

func (g *GraphList) V() int {
	return g.v
}

func (g *GraphList) Adjacency(v int) []int {
	if v < 0 || v >= g.v {
		panic("vertex v invalid")
	}

	l := g.g[v]
	adjacent := make([]int, 0, g.v)
	for e := l.Front(); e != nil; e = e.Next() {
		adjacent = append(adjacent, e.Value.(int))
	}

	return adjacent
}

func (g *GraphList) String() string {
	var bf bytes.Buffer

	bf.WriteString(fmt.Sprintf("Graph V: %d\n", g.v))
	// adjancet edge
	for i := 0; i < g.v; i++ {
		bf.WriteString(fmt.Sprintf("Vertex %d: ", i))
		adjacent := g.Adjacency(i)
		bf.WriteString(fmt.Sprintf("%v\n", adjacent))
	}

	return bf.String()
}
