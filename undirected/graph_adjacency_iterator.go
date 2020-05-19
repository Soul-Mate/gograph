package undirected

import "container/list"

type GraphAdjacencyIterator interface {
	Next() int
	HasNext() bool
}

type GraphMatrixAdjacencyIterator struct {
	i int
	n int
	v []int
}

func NewGraphMatrixAdjacencyIterator(v []int) *GraphMatrixAdjacencyIterator {
	return &GraphMatrixAdjacencyIterator{
		0, len(v), v,
	}
}

func (g *GraphMatrixAdjacencyIterator) Next() int {
	v := g.v[g.i]
	g.i++
	return v
}

func (g *GraphMatrixAdjacencyIterator) HasNext() bool {
	return g.i == g.n
}

type GraphListAdjacencyIterator struct {
	e *list.Element
}

func NewGraphListAdjacencyIterator(l *list.List) *GraphListAdjacencyIterator {
	return &GraphListAdjacencyIterator{
		l.Front(),
	}
}

func (g *GraphListAdjacencyIterator) Next() int {
	v := g.e.Value.(int)
	g.e = g.e.Next()
	return v
}

func (g *GraphListAdjacencyIterator) HasNext() bool {
	return g.e == nil
}
