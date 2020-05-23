package undirected

type BipartiteDetection struct {
	g           Graph
	visited     []bool
	colors      []int
	isBipartite bool
}

func NewBipartiteDetection(g Graph) *BipartiteDetection {
	b := new(BipartiteDetection)
	b.g = g
	b.isBipartite = true
	b.visited = make([]bool, b.g.V())
	b.colors = make([]int, b.g.V())
	for i := 0; i < b.g.V(); i++ {
		b.visited[i] = false
		b.colors[i] = -1
	}

	for i := 0; i < b.g.V(); i++ {
		if !b.visited[i] {
			if !b.dfs(i, 0) {
				b.isBipartite = false
				break
			}
		}
	}

	return b
}

func (b *BipartiteDetection) dfs(v, color int) bool {
	b.visited[v] = true
	b.colors[v] = color
	it := b.g.Adjacency(v)

	for !it.HasNext() {
		w := it.Next()
		if !b.visited[w] {
			if !b.dfs(w, 1-color) {
				return false
			}

		} else if b.colors[v] == b.colors[w] {
			return false
		}
	}

	return true
}

func (b *BipartiteDetection) IsBipartite() bool {
	return b.isBipartite
}
