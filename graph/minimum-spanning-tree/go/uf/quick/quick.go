package quick

// 算法一：quick-find
type QuickFind struct {
	count int   // 分量数量
	id    []int // 分量值
}

func New(n int) *QuickFind {
	qf := &QuickFind{
		count: n,
		id:    make([]int, n),
	}

	// init node's id to itself
	for i := 0; i < n; i++ {
		qf.id[i] = i
	}

	return qf
}

func (qf *QuickFind) Count() int {
	return qf.count
}

func (qf *QuickFind) Connected(p, q int) bool {
	if p >= len(qf.id) || q >= len(qf.id) {
		return false
	}

	return qf.id[p] == qf.id[q]
}

func (qf *QuickFind) Find(v int) int {
	if v >= len(qf.id) {
		return -1
	}
	return qf.id[v]
}

func (qf *QuickFind) Union(p, q int) {
	pid := qf.Find(p)
	qid := qf.Find(q)

	if pid == qid {
		return
	}

	// 遍历所有节点，将分量为pid的节点重命名为qid
	for i, v := range qf.id {
		if v == pid {
			qf.id[i] = qid
		}
	}

	qf.count--
}
