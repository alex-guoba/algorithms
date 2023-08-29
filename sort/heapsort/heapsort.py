
"""
    快速排序
    参考：https://zq99299.github.io/dsalg-tutorial/dsalg-java-hsp/11/01.html#%E6%AD%A5%E9%AA%A4%E4%BA%8C-%E5%B0%86%E5%A0%86%E9%A1%B6%E5%85%83%E7%B4%A0%E4%B8%8E%E6%9C%AB%E5%B0%BE%E5%85%83%E7%B4%A0%E8%BF%9B%E8%A1%8C%E4%BA%A4%E6%8D%A2

"""

from typing import TypeAlias

#Vector = list[int]
Vector: TypeAlias = list[int]

def max_heapify(src: Vector, index: int, size: int) -> Vector:
    left = 2*index + 1
    right = 2*index + 2

    current = index
    if left < size and src[left] > src[current]:
        current = left
    if right < size and src[right] > src[current]:
        current = right

    if current != index:
        src[index], src[current] = src[current], src[index]
        max_heapify(src, current, size)

def heapsort(src: Vector) -> Vector:
    if len(src) <= 1:
        return src

    # 1. construct a max-heap from a array. all leaf node shift-up
    size = len(src)
    # max_heapify from right to left but exclude leaves (last level)
    for i in range(size // 2 - 1, -1, -1):
        max_heapify(src, i, size)

    # 2. move header node to the tail and re-heapify the next.
    for i in range(size-1, -1, -1):
        src[0], src[i] = src[i], src[0]
        max_heapify(src, 0, i)

    return src




