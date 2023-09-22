package uf

// union-find算法，用于判断两个基点是否已经有连接关系
//
// 《算法》中给出了三种解法：quick-find、union-find以及加权union-find
// 整体思路：
// - 为每个node节点分配一个所属分组（分量）id，初始时每个node的分组id都不相同。
// - union(a, b): 每次连接两个（未连接的）节点时，将其分量id设置为同一个分量id即可
// - connected(a, b)：判断两个节点是否已联通，只需要判断其分量id是否相等即可。
//
// union-find与加权union实现上差别不大，此处仅实现第三种

type UF interface {
	// 连接两个节点
	Union(q int, p int)

	// 查找节点所属分量
	Find(v int) int

	// 是否已连接
	Connected(p int, q int) bool

	// 联通分量数量
	Count() int
}
