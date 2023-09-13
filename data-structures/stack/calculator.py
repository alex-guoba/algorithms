""""
E.W.Dijkstra在20世纪60年代发明了一个非常简单的算法，用两个栈（一个用于保存运算符，一个用于保存操作数）完成简单的数值计算。
每当遇到右括号时，从运算符堆栈中取出最近插入的运算法，然后从操作数堆栈中取出对应的操作数并进行计算操作。计算完毕后将结果append到堆栈。循环反复，在处理完最后一个右括号之后，操作数栈上只会有一个值
"""


from collections import deque
from enum import Enum

class TokenType(Enum):
    Num = 1
    Flag = 2
    RB = 3

class Token:
    def __init__(self, token_type, val):
        self.token_type = token_type
        self.val = val

    def __str__(self):
        return 'Type: %s, Val: %s' % (self.token_type, str(self.val))

class TokenParser:
    CFlags = ('+', '-', '*', '/')
    
    @classmethod
    def parse_tokens(cls, statement: str) -> list[Token]:
        left: str = statement
        tokens: list[Token] = []

        while(True):
            left, token = cls.next_token(left)
            if not token:
                break
            else:
                tokens.append(token)
        return tokens

    @classmethod
    def next_token(cls, left: str) -> (str, Token | None):
        if not left:
            return "", None
        ch = left[0]
        left = left[1:]

        if ch in cls.CFlags:
            return left, Token(TokenType.Flag, ch)
        if ch.isdigit():
            return left, Token(TokenType.Num, int(ch))
        if ch == ')':
            return left, Token(TokenType.RB, '')

        return cls.next_token(left)


class Calculator:
    @classmethod
    def calc_once(cls, numbers: deque, flags: deque):
        # TODO: check empty
        flag = flags.pop()
        match flag.val:
            case '+':
                second = numbers.pop()
                first = numbers.pop()
                numbers.append(Token(TokenType.Num, first.val + second.val))

            case '-':
                second = numbers.pop()
                first = numbers.pop()
                numbers.append(Token(TokenType.Num, first.val - second.val))

            case '*':
                second = numbers.pop()
                first = numbers.pop()
                numbers.append(Token(TokenType.Num, first.val * second.val))

            case '/':
                second = numbers.pop()
                first = numbers.pop()
                numbers.append(Token(TokenType.Num, int(first.val / second.val)))

            case '_':
                print('invalid')

    @classmethod
    def calc(cls, statement: str) -> int:
        tokens = TokenParser.parse_tokens(statement)
        numbers = deque()
        flags = deque()

        for token in tokens:
            # print("token: %s" % str(token))
            if token.token_type == TokenType.Num:
                numbers.append(token)
            elif token.token_type == TokenType.Flag:
                flags.append(token)
            else:
                # right bracket, pop up the number and flag to calulate
                cls.calc_once(numbers, flags)

        if len(flags):
            cls.calc_once(numbers, flags)
        number = numbers.pop()
        return number.val

if __name__ == '__main__':
    print(Calculator.calc('1+((2+3)*(4*5))'))
