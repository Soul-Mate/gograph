package undirected

type CycleDetection struct {
	g       Graph
	visited []bool
}

func NewCycleDetection(g Graph) *CycleDetection {
	c := new(CycleDetection)
	c.g = g
	c.visited = make([]bool, c.g.V())
	return c
}

func (c *CycleDetection) Detection(s int) bool {
	for i := 0; i < c.g.V(); i++ {
		if c.dfs(s, s) {
			return true
		}
	}

	return false
}

func (c *CycleDetection) dfs(v, p int) bool {
	c.visited[v] = true
	it := c.g.Adjacency(v)
	for !it.HasNext() {
		w := it.Next()
		if !c.visited[w] {
			if c.dfs(w, v) {
				return true
			}

		} else if w != p {
			return true
		}
	}

	return false
}
