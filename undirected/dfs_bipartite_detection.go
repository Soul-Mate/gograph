package undirected

type BipartiteDetection struct {
	g           Graph
	visited     []bool
	isBipartite bool
}

func NewBipartiteDetection(g Graph) *BipartiteDetection {
	b := new(BipartiteDetection)
	b.g = g
	b.isBipartite = true
	b.visited = make([]bool, b.g.V())

	for i := 0; i < b.g.V(); i++ {
		if !b.visited[i] && b.dfs(i) {
			b.isBipartite = true
			break
		}
	}

	return b
}
