package gograph

type Graph interface {
	V() int
	Adjacency(v int) GraphAdjacencyIterator
	IsConnected(int, int) bool
}
