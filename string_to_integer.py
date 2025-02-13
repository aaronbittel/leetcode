class Solution:
    def myAtoi(self, s: str) -> int:
        int_str = ""
        is_pos = True
        started = False
        break_loop = False
        read_zero = False
        for ch in s:
            match ch:
                case " ":
                    if read_zero and not started:
                        return 0
                    if started:
                        break_loop = True
                case "-":
                    if read_zero and not started:
                        return 0
                    if not started:
                        started = True
                        is_pos = False
                    else:
                        break_loop = True
                case "+":
                    if read_zero and not started:
                        return 0
                    if not started:
                        started = True
                        is_pos = True
                    else:
                        break_loop = True
                case "0":
                    read_zero = True
                    if started and int_str:
                        int_str += "0"
                case n if n in "123456789":
                    if not started:
                        started = True
                    int_str += n
                case _:
                    break_loop = True

            if break_loop:
                break

        res = 0
        digit = 1
        for n in int_str[::-1]:
            res += digit * int(n)
            digit *= 10

        if not is_pos:
            return max(-res, -(2**31))
        return min(res, 2**31 - 1)


from collections import namedtuple  # noqa: E402

tests = namedtuple("Test", ["ins", "out"])
tests = [
    tests(ins="42", out=42),
    tests(ins="-042", out=-42),
    tests(ins="1337c0d3", out=1337),
    tests(ins="0-1", out=0),
    tests(ins="words and 987", out=0),
    tests(ins="0  123", out=0),
    tests(ins="  0000000000012345678", out=12345678),
]

for tt in tests:
    got = Solution().myAtoi(tt.ins)
    if got != tt.out:
        print(f"Ins: {tt.ins} -> Expected: {tt.out}, got: {got}")
    else:
        print(f"Ins: {tt.ins} -> {got}: Success!")
