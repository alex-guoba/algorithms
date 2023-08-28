"""
给定含有N个互不相同元素的序列，输出这N个元素的所有排列（全排列）。

以字符串"abc"为例，全排列包括"abc","acb","bac","bca","cab","cba"这6个结果。

解法1：使用状态标识位记录每个节点的访问状态，dfs搜索过程中，跳过已经访问过的节点，
解法2：使用queue记录还未访问的节点。回退时弹出的队首节点插入队尾，循环使用
"""

from typing import TypeAlias

#Vector = list[int]
Vector: TypeAlias = list[int | str]


def all_permutations(sequence: Vector) -> None:
    current: Vector = []
    dfs(sequence, current)
    pass

def dfs(
        sequence: Vector,
        current: Vector
        ) -> None:

    if 0 == len(sequence):
        print(current)
        return

    cycle = len(sequence)
    for i in range(cycle):
        current.append(sequence[0])
        sequence = sequence[1:] # pop head

        dfs(sequence, current)

        tail = current.pop()
        sequence.append(tail)


if __name__ == "__main__":
    all_permutations([1, 2, 3, 4])

