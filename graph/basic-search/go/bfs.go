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
	"sort"
)

func BFS(graph *Graph, vertice int) []int {
	visited := map[int]bool{}

	queue := []int{vertice}
	routes := []int{}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:] // queue pop

		if _, ok := visited[v]; ok {
			continue
		}

		// DO visit: mark & routes
		visited[v] = true
		routes = append(routes, v) // append to routes

		if adjs, ok := graph.GetAdjacent(v); ok {
			// sort, make sure routes consistent for unit test
			sort.Ints(adjs)

			for _, adj := range adjs {
				// only unvisited node will be append
				if _, ok := visited[adj]; ok {
					continue
				}

				queue = append(queue, adj)
			}
		}
	}
	return routes
}
