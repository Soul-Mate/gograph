package gograph

type GraphVertexT interface{}

type GraphEdgeT struct {
	V1 GraphVertexT
	V2 GraphVertexT
}

type Vertex struct {
	Lable interface{}
}
