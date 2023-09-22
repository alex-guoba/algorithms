package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alex-guoba/mst/uf/quick"
	"github.com/alex-guoba/mst/uf/union"
)

func newGraph() *Graph {
	g := New(7)

	for i := 0; i < 7; i++ {
		g.AddVertex(i)
	}

	g.AddWeightedEdge(0, 1, 4)
	g.AddWeightedEdge(0, 2, 3)
	g.AddWeightedEdge(0, 5, 9)
	g.AddWeightedEdge(1, 2, 10)
	g.AddWeightedEdge(2, 3, 7)
	g.AddWeightedEdge(2, 4, 2)
	g.AddWeightedEdge(3, 4, 5)
	g.AddWeightedEdge(3, 5, 11)
	g.AddWeightedEdge(4, 6, 1)
	g.AddWeightedEdge(6, 1, 6)
	g.AddWeightedEdge(6, 5, 13)

	return g
}

func TestKruskal(t *testing.T) {
	type args struct {
		graph *Graph
	}

	// construt graph
	graph := newGraph()

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"kruskal",
			args{graph},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qkfinder := quick.New(graph.Vertices)
			kr := Kruskal(tt.args.graph, qkfinder)

			assert.Equal(t, qkfinder.Count(), 1, "all nodes connected")

			uffinder := union.New(graph.Vertices)
			ur := Kruskal(tt.args.graph, uffinder)
			assert.Equal(t, uffinder.Count(), 1, "all nodes connected")

			assert.Equal(t, kr, ur, "finder exception")
		})
	}
}
