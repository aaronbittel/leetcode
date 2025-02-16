from defs import test
import logging

logging.basicConfig(format="", level=logging.INFO)


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        if s == "":
            return 0

        open, close, res = 0, 0, 0

        for ch in s:
            if ch == "(":
                open += 1
            else:
                close += 1

            if open == close:
                res = max(res, open * 2)
            elif close > open:
                open, close = 0, 0

        open, close = 0, 0
        for ch in s[::-1]:
            if ch == "(":
                open += 1
            else:
                close += 1

            if open == close:
                res = max(res, open * 2)
            elif close < open:
                open, close = 0, 0

        return res

    TESTS = [
        test(ins="(()", want=2),
        test(ins=")()())", want=4),
        test(ins="", want=0),
        test(ins="((((())))", want=8),
    ]
