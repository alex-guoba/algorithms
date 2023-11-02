"""
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

输入：n = 1
输出：["()"]


思路：
1. 判断左右括号所剩的数量，最初始都是n;当左括号（(）有剩余，继续做选择；
2. 判断当右括号比左括号剩的多，才能选右括号；继续递归做选择
3. 出口:构建的字符串是 2n的时候，此时已经该分支已经构建完成，加入选项；
"""

def generate(left: str, right: str, result: str):
    # stop
    if left == 0 and right == 0:
        print(result)
        return

    if left >= 0:
        generate(left-1, right, result+"(")

    if left < right:
        generate(left, right-1, result+")")

def generate_brace(n: int):
    result = ""
    generate(n, n, result)

if __name__ == "__main__":
    generate_brace(3)
