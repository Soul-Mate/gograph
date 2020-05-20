package undirected

type DFS struct {
	g       Graph
	order   []int
	visited []bool
}

func NewDFS(g Graph) *DFS {
	dfs := new(DFS)
	dfs.g = g
	dfs.visited = make([]bool, dfs.g.V())

	for i, n := 0, dfs.g.V(); i < n; i++ {
		dfs.visited[i] = false
	}

	return dfs
}

func (dfs *DFS) DFS() {
	for i, vn := 0, dfs.g.V(); i < vn; i++ {
		if !dfs.visited[i] {
			dfs.dfs1(i)
		}
	}
}

func (dfs *DFS) dfs1(v int) {
	dfs.visited[v] = true
	dfs.order = append(dfs.order, v)
	it := dfs.g.Adjacency(v)
	for !it.HasNext() {
		w := it.Next()
		if !dfs.visited[w] {
			dfs.dfs1(w)
		}
	}
}

func (dfs *DFS) Order() []int {
	return dfs.order
}
