from defs import tests


class Solution:
    TESTS = [
        tests(ins=("11", "1"), out="100"),
        tests(ins=("1010", "1011"), out="10101"),
    ]

    def addBinary(self, a: str, b: str) -> str:
        res = ""

        if len(a) < len(b):
            a, b = b, a
        b = "0" * (len(a) - len(b)) + b

        carryover = 0

        for n1, n2 in zip(a[::-1], b[::-1]):
            match (n1, n2, carryover):
                case ("1", "1", 0):
                    res += "0"
                    carryover = 1
                case ("1", "1", 1):
                    res += "1"
                    carryover = 1
                case ("1", "0", 0) | ("0", "1", 0):
                    res += "1"
                    carryover = 0
                case ("1", "0", 1) | ("0", "1", 1):
                    res += "0"
                    carryover = 1
                case ("0", "0", 0):
                    res += "0"
                    carryover = 0
                case ("0", "0", 1):
                    res += "1"
                    carryover = 0

        if carryover == 1:
            res += "1"

        return res[::-1]
