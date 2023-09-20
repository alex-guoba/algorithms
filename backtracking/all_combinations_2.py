"""
>>> generate_all_combinations(n=4, k=2)
        [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]

mark:  未实现剪枝优化
"""
Result = list[list[int]]

def generate_all_combinations(n: int, k: int):
    result: Result = []

    backtrack(result, 1, n, k, [])
    return result

def backtrack(
        result: Result,
        start: int,
        end: int,
        level: int,
        current: []
        ):
    if level == 0:
        result.append(current[:])
        print("append current %s, result, %s" % (current, result))
        return

    for i in range(start, end+1):
        current.append(i)
        backtrack(result, i+1, end, level - 1, current)
        current.pop()

def print_all_state(total_list: Result) -> None:
    for i in total_list:
        print(*i)

if __name__ == "__main__":
      n = 4
      k = 2
      total_list = generate_all_combinations(n, k)
      print_all_state(total_list)
