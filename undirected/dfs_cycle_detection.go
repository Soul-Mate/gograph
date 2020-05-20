package undirected

type CycleDetection struct {
	g        Graph
	visited  []bool
	hasCycle bool
}

func NewCycleDetection(g Graph) *CycleDetection {
	c := new(CycleDetection)
	c.g = g
	c.visited = make([]bool, c.g.V())

	for i := 0; i < c.g.V(); i++ {
		if !c.visited[i] && c.dfs(i, i) {
			c.hasCycle = true
			break
		}
	}

	return c
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

func (c *CycleDetection) HasCycle() bool {
	return c.hasCycle
}
