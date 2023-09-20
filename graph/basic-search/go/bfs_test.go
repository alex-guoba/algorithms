// BFS定义及实现
//
// 定义：单点最短路径。给定一幅图和一个起点s，回答“从s到给定目的顶点v是否存在一条路径？
// 如果有，找出其中最短的那条（所含边数最少）。”等类似问题。 解决这个问题的经典方法叫做广度优先搜索（BFS）
//
// 算法：使用一个队列来保存所有已经被标记过但其邻接表还未被检查过的顶点。先将起点加入队列，然后重复以下步骤直到队列为空：
// 	- 取队列中的下一个顶点v并标记它；
// 	- 将与v相邻的所有未被标记过的顶点加入队列。

package graph

import (
	"reflect"
	"testing"
)

func newGraph(vertice int) *Graph {
	g := New(vertice)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	return g
}

func TestBFS(t *testing.T) {
	type args struct {
		graph   *Graph
		vertice int
	}

	// construt graph
	graph := newGraph(0)

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"basic 1",
			args{
				graph,
				0,
			},
			[]int{0, 1, 2, 5, 3, 4},
		},
		{
			"basic 2",
			args{
				graph,
				4,
			},
			[]int{4, 2, 3, 0, 1, 5},
		},
		{
			"basic 3",
			args{
				graph,
				1,
			},
			[]int{1, 0, 2, 5, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.graph, tt.args.vertice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
