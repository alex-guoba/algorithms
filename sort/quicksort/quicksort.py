
"""
    快速排序

"""

from typing import TypeAlias

#Vector = list[int]
Vector: TypeAlias = list[int]

def quicksort(src: Vector) -> Vector:
    if len(src) <= 1:
        return src

    pivot = src[0]

    left :Vector = list(filter(lambda x: x < pivot, src))
    right :Vector = list(filter(lambda x: x > pivot, src))

    return quicksort(left) + [pivot] + quicksort(right)
