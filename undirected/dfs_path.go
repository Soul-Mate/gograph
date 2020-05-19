package undirected

type Path struct {
	g       Graph
	src     int
	pre     []int
	visited []bool
}

func NewPath(g Graph, s int) *Path {
	p := new(Path)
	p.g = g
	p.src = s
	p.pre = make([]int, p.g.V())
	p.visited = make([]bool, p.g.V())
	for i := 0; i < p.g.V(); i++ {
		p.visited[i] = false
		p.pre[i] = -1
	}

	return p
}

func (p *Path) Path(t int) []int {
	if !p.dfs(p.src, t, p.src) {
		return nil
	}

	path := make([]int, 0)
	for cur := p.pre[t]; cur != p.src; cur = p.pre[cur] {
		path = append(path, cur)
	}

	path = append(path, p.src)
	return path
}

func (p *Path) dfs(v, t, pre int) bool {
	p.pre[v] = pre
	p.visited[v] = true
	if v == t {
		return true
	}

	it := p.g.Adjacency(v)
	for !it.HasNext() {
		w := it.Next()
		if !p.visited[w] {
			if p.dfs(w, t, v) {
				return true
			}
		}
	}

	return false
}
