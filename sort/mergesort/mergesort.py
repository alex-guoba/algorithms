from typing import TypeAlias

#Vector = list[int]
Vector: TypeAlias = list[int]

def mergesort(src: Vector) -> Vector:
    if len(src) <= 1:
        return src
    
    mid = int(len(src)/2)
    left = mergesort(src[:mid])
    right = mergesort(src[mid:])

    dst: Vector = []

    while(True):
        if not left:
            dst.extend(right)
            break

        if not right:
            dst.extend(left)
            break

        if left[0] < right[0]:
            dst.append(left[0])
            left = left[1:]
        else:
            dst.append(right[0])
            right = right[1:]

    return dst


if __name__ == '__main__':
    print(mergesort([5,8,2,7,4]))




