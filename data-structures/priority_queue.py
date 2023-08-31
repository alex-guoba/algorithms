"""
  基础优先队列，使用列表实现，不考虑排序、堆实现
  复杂度:
  - pop: O(1)
  - enqueue: O(N)
"""

from typing import TypeAlias, Any

ItemAlias: TypeAlias = Any

class Node:
    """node type"""
    def __init__(self, value: ItemAlias, priority: int) -> None:
        self.value = value
        self.priority = priority
        self.next = None

class PriorQueue:
    """
    优先队列
    """
    def __init__(self) -> None:
        self.head = None
        self.size = 0

    def is_empty(self) -> bool:
        """
        判断队列是否为空
        """
        return self.size == 0

    def enqueue(self, value: ItemAlias, priority: int) -> None:
        """
        插入元素
        """
        node = Node(value, priority)
        if self.is_empty():
            self.head = node
        elif priority > self.head.priority:
            node.next = self.head
            self.head = node # reset head
        else:
            cur = self.head
            while cur.next and cur.next.priority >= priority:
                cur = cur.next
            node.next = cur.next
            cur.next = node
        self.size = self.size + 1

    def pop(self) -> Node:
        """
        弹出元素
        """
        if self.is_empty():
            return None
        header = self.head
        self.head = self.head.next
        self.size = self.size - 1
        return header

if __name__ == '__main__':
    pq = PriorQueue()
    pq.enqueue(1, 1)
    pq.enqueue(2, 2)
    pq.enqueue(3, 3)
    pq.enqueue(4, 4)
    pq.enqueue(5, 5)
    pq.enqueue(6, 6)
    pq.enqueue(7, 7)
    pq.enqueue(8, 8)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    print(pq.pop().value)
    assert(pq.is_empty() is True)

