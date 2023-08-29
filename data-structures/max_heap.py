"""
最大堆实现:
    - 节点数据任意类型
    - 支持peek、push
"""

class Node:
    def __init__(self, name, val):
        self.name = name
        self.val = val

    def __str__(self):
        return f"{self.__class__.__name__}({self.name}, {self.val})"

    def __lt__(self, other):
        return self.val < other.val


class MaxHeap:

    def __init__(self, nodelists=None):
        self.nodes = []
        for node in nodelists:
            self.push(node)

    def peek(self) -> Node:
        if not self.nodes:
            return

        tail = self.nodes.pop()
        if not self.nodes: # last one
            return tail

        node = self.nodes[0]
        self.nodes[0] = tail
        self._shift_down(0)
        return node

    def push(self, node: Node):
        self.nodes.append(node) # append to tail
        self._shift_up()

    def dump(self):
        print('-'*10)
        for node in self.nodes:
            print(str(node))
        print('-'*10)

    def _parent(self, idx: int) -> int | None:
        if idx == 0:
            return None
        return int((idx-1)/2)

    def _left(self, idx: int) -> int | None:
        if idx*2+1 < len(self.nodes):
            return idx*2+1
        return None

    def _right(self, idx: int) -> int | None:
        if idx*2+2 < len(self.nodes):
            return idx*2+2
        return None

    def _max_node(self, left: int | None, right: int | None) -> int | None:
        if left is not None and right is not None:
            _l = self.nodes[left]
            _r = self.nodes[right]
            if _l > _r: # chose bigger one
                return left
            else:
                return right

        return left or right

    def _shift_down(self, idx):
        while(idx < len(self.nodes)):
            curr = self.nodes[idx]

            left = self._left(idx)
            right = self._right(idx)

            child = self._max_node(left, right)
            if not child:
                break

            # compare and exchange
            if curr > self.nodes[child]:
                break
            self.nodes[idx] = self.nodes[child]
            self.nodes[child] = curr

            idx = child # continue

    def _shift_up(self):
        if not self.nodes:
            return

        idx_cur = len(self.nodes) - 1
        idx_parent = self._parent(idx_cur)
        while idx_parent is not None:
            node_cur, node_parent = self.nodes[idx_cur], self.nodes[idx_parent]

            if node_parent > node_cur:
                break

            # exchange position
            self.nodes[idx_cur], self.nodes[idx_parent] = node_parent, node_cur

            idx_cur = idx_parent
            idx_parent = self._parent(idx_cur)


if __name__ == '__main__':
    #import doctest
    #doctest.testmod()

    r = Node("R", -1)
    b = Node("B", 6)
    a = Node("A", 3)
    x = Node("X", 1)
    e = Node("E", 4)
    m = Node("M", 2)

    mheap = MaxHeap([r, b, a, x, e, m])

    #mheap.dump()

    print(mheap.peek())
    print(mheap.peek())
    print(mheap.peek())
    print(mheap.peek())
    print(mheap.peek())
    print(mheap.peek())
    print(mheap.peek())
