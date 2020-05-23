package undirected

type SingleSourcePath struct {
	g       Graph
	src     int
	pre     []int
	visited []bool
}

func NewSingleSourcePath(g Graph, src int) *SingleSourcePath {
	s := new(SingleSourcePath)
	s.g = g
	s.src = src
	s.pre = make([]int, s.g.V())
	s.visited = make([]bool, s.g.V())

	for i, n := 0, s.g.V(); i < n; i++ {
		s.visited[i] = false
		s.pre[i] = -1
	}

	s.dfs(s.src, s.src)
	return s
}

func (s *SingleSourcePath) Path(t int) []int {
	if !s.IsConnected(t) {
		return nil
	}

	path := make([]int, 0, s.g.V())
	for cur := s.pre[t]; cur != s.src; cur = s.pre[cur] {
		path = append(path, cur)
	}

	path = append(path, s.src)

	return path
}

func (s *SingleSourcePath) IsConnected(t int) bool {
	return s.visited[t]
}

func (s *SingleSourcePath) dfs(v, p int) {
	s.visited[v] = true
	s.pre[v] = p
	it := s.g.Adjacency(v)
	for !it.HasNext() {
		w := it.Next()
		if !s.visited[w] {
			s.dfs(w, v)
		}
	}
}
