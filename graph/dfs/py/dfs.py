from collections import defaultdict
from collections import deque
from queue import Queue

class Graph:

    def __init__(self) -> None:
        self.nodes = defaultdict(list)

    def add_edge(self, vfrom: int, vto: int):
        self.nodes[vfrom].append(vto)
        self.nodes[vto].append(vfrom)

    def adj(self, v):
        return self.nodes[v]

    def print(self):
        for v, edges in self.nodes.items():
            es = map(lambda x: str(x), edges)
            print(str(v),  " : ",  " -> ".join(es))

    def dfs(self, vhead):
        """
        无向图DFS实现：
        使用栈的FILO的特定，每次先沿某条边遍历到无法继续为止，然后再弹出最近插入的节点，继续执行遍历。
        """
        visited = {}
        stack = deque()

        stack.append(vhead) # push right
        while stack:
            v = stack.pop() # pop right
            if visited.get(v):
                continue

            visited[v] = True
            print(v)
            adjs = self.adj(v)
            for n in adjs:
                if visited.get(n):
                    continue
                stack.append(n)

    def bfs(self, vhead):
        """
        无向图BFS实现
        利用队列
        """
        visited = {}
        queue = Queue() # FIFO queue
        
        queue.put(vhead)
        while not queue.empty():
            v = queue.get()
            if visited.get(v):
                continue

            visited[v] = True
            print(v)
            adjs = self.adj(v)
            for n in adjs:
                if visited.get(n):
                    continue
                queue.put(n)
        
if __name__ == '__main__':
    d = Graph()
    d.add_edge(0, 5)
    d.add_edge(4, 3)
    d.add_edge(2, 1)
    d.add_edge(9, 12)
    d.add_edge(6, 4)
    d.add_edge(5, 4)
    d.add_edge(0, 2)
    d.add_edge(11, 12)
    d.add_edge(9, 10)
    d.add_edge(0, 6)
    d.add_edge(7, 8)
    d.add_edge(9, 11)
    d.add_edge(5, 3)
    #d.print()

    d.bfs(5)