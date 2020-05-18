package gograph

type DFS struct {
	g       Graph
	visited []int
}

func NewDFS(g Graph) *DFS {
	dfs := new(DFS)
	dfs.g = g
	dfs.visited = make([]int, dfs.g.V())

	for i, n := 0, dfs.g.V(); i < n; i++ {
		dfs.visited[i] = -1
	}

	return dfs
}

func (dfs *DFS) DFS() {

}
