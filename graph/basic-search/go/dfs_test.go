// BFS定义及实现
//
// 定义：中文名是深度优先搜索，是一种用于遍历或搜索树或图的算法。所谓深度优先，就是说每次都尝试向更深的节点走。
// 	该算法讲解时常常与 BFS 并列，但两者除了都能遍历图的连通块以外，用途完全不同，很少有能混用两种算法的情况。
//
// 算法：与bfs不同的是利用栈的LIFO的特点来遍历，这样每次最后插入的节点会先弹出来达到优先向更深处遍历的效果

package graph

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	type args struct {
		graph   *Graph
		vertice int
	}

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
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"basic 2",
			args{
				graph,
				4,
			},
			[]int{4, 2, 0, 1, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DFS(tt.args.graph, tt.args.vertice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
