// 最小生成树算法中的kruskal算法
//
// 定义：图的生成树是它的一棵含有其所有顶点的无环连通子图。
//  一幅加权图的最小生成树（MST）是它的一棵权值（树中所有边的权值之和）最小的生成树。
// 两个特性：
// - 用一条边连接树中的任意两个顶点都会产生一个新的环；
// - 从树中删去一条边将会得到两棵独立的树。
//
// kruskal算法：按照边的权重顺序（从小到大）处理它们，将边加入最小生成树中，
// 加入的边不会与已经加入的边构成环，直到树中含有V-1条边为止。
// 这些边逐渐由一片森林合并为一棵树，也就是图的最小生成树。
//
// 证明:如果下一条将被加入最小生成树中的边不会和已有的边构成环，那么它就跨越了由所有和树顶点相邻的顶点组成的集合以及它们的补集所构成的一个切分。
// 因为加入的这条边不会形成环、它是目前已知的唯一一条横切边且是按照权重顺序选择的边，所以它必然是权重最小的横切边。因此，该算法能够连续选择权重最小的横切边，和贪心算法一致。

package graph

import (
	"sort"

	"github.com/alex-guoba/mst/uf"
)

func Kruskal(graph *Graph, finder uf.UF) [][]int {
	edges := graph.GetEdges()
	msf := [][]int{}

	// sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// iterate edges
	for _, edge := range edges {
		p, q, weight := edge[0], edge[1], edge[2]

		if finder.Connected(p, q) {
			continue
		}

		finder.Union(p, q)
		msf = append(msf, []int{p, q, weight})

		// full, stop iteration
		if len(msf) == graph.Vertices-1 {
			break
		}
	}

	return msf
}
