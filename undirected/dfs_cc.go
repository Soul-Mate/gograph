package undirected

type CC struct {
	g       Graph
	ccount  int // connected component count
	visited []int
}

func NewCC(g Graph) *CC {
	cc := new(CC)
	cc.g = g
	cc.ccount = 0
	cc.visited = make([]int, cc.g.V())

	for i, n := 0, cc.g.V(); i < n; i++ {
		cc.visited[i] = -1
	}

	for i, n := 0, cc.g.V(); i < n; i++ {
		if cc.visited[i] == -1 {
			cc.dfs(i)
			cc.ccount++
		}
	}

	return cc
}

func (cc *CC) ConnectedComponentCount() int {
	return cc.ccount
}

func (cc *CC) ConnectedComponent(v int) []int {
	if v < 0 || v >= cc.g.V() {
		panic("vertex v invalid")
	}

	res := make([]int, 0, cc.ccount)
	ccount := cc.visited[v]
	if ccount == -1 {
		return res
	}

	for i := 0; i < len(cc.visited); i++ {
		if cc.visited[i] == ccount {
			res = append(res, i)
		}
	}

	return res
}

func (cc *CC) dfs(v int) {
	cc.visited[v] = cc.ccount
	it := cc.g.Adjacency(v)
	for !it.HasNext() {
		w := it.Next()
		if cc.visited[w] == -1 {
			cc.dfs(w)
		}
	}
}
