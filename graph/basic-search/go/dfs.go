// BFS定义及实现
//
// 定义：中文名是深度优先搜索，是一种用于遍历或搜索树或图的算法。所谓深度优先，就是说每次都尝试向更深的节点走。
// 	该算法讲解时常常与 BFS 并列，但两者除了都能遍历图的连通块以外，用途完全不同，很少有能混用两种算法的情况。
//
// 算法：与bfs不同的是利用栈的LIFO的特点来遍历，这样每次最后插入的节点会先弹出来达到优先向更深处遍历的效果

package graph

import (
	"sort"
)

func DFS(graph *Graph, vertice int) []int {
	visited := map[int]bool{}

	queue := []int{vertice}
	routes := []int{}

	for len(queue) > 0 {
		// stack pop
		v := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if _, ok := visited[v]; ok {
			continue
		}

		// DO visit
		visited[v] = true
		routes = append(routes, v) // append to routes

		if adjs, ok := graph.GetAdjacent(v); ok {
			// sort, make sure routes consistent for unit test
			sort.Sort(sort.Reverse(sort.IntSlice(adjs)))

			for _, adj := range adjs {
				if _, ok := visited[adj]; ok {
					continue
				}

				//visited[adj] = true
				//routes = append(routes, adj) // append to routes
				queue = append(queue, adj)
			}
		}
	}
	return routes
}
