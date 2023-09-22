package union

// 算法说明
// quick-find的问题在于每次union都要遍历所有节点，效率太低
// union-find的思路是将同一个分量的节点构造成一颗树，精确的分量id值都存储在该分量树的root节点中：
// - 同一分量中任意两个节点的root节点是相同的，所以分量值也一定相同
// - 对于非root节点，分量值存储的其实是指向parent节点的索引（目的等同于指针，便于向上回溯）
//
// union-find有个潜在的风险，极端情况下树的高度等于节点数，这样会影响Find查找效率。
// 为解决此问题，加权uinon-find引入了分量大小的概念（可以理解为数的高度），每次合并时优先将较小的树合并到大的树，有效控制数的高度

// 算法三：加权union-find
type UnionFind struct {
	count int // 分量数量

	id   []int // 分量值
	rank []int // 分量的大小，初始值都为1，每次union时将较小的数合并到加到的数，保证数的高度小于未加权的版本
}

func New(n int) *UnionFind {
	uf := &UnionFind{
		count: n,
		id:    make([]int, n),
		rank:  make([]int, n),
	}

	// init node's id to itself
	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.rank[i] = 1
	}

	return uf
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) Connected(p, q int) bool {
	if p >= len(uf.id) || q >= len(uf.id) {
		return false
	}

	return uf.id[p] == uf.id[q]
}

func (uf *UnionFind) Find(v int) int {
	if v >= len(uf.id) {
		return -1
	}

	// root
	if uf.id[v] == v {
		return v
	}
	return uf.Find(uf.id[v])
}

// 分量树重的深度
func (uf *UnionFind) Depth(v int, current int) int {
	// root
	if uf.id[v] == v {
		return current
	}
	return uf.Depth(uf.id[v], current+1)
}

func (uf *UnionFind) Union(p, q int) {
	pid := uf.Find(p) // 默认p节点合入q节点
	qid := uf.Find(q)

	if pid == qid {
		return
	}

	prank := uf.rank[pid]
	qrank := uf.rank[qid]

	// switch
	if prank > qrank {
		pid, qid = qid, pid
	}

	uf.id[pid] = qid
	uf.rank[qid] += uf.rank[pid]

	uf.count--
}
