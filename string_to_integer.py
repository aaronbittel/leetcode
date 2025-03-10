from defs import test


class Solution:
    TESTS = [
        test(ins="42", want=42),
        test(ins="-042", want=-42),
        test(ins="1337c0d3", want=1337),
        test(ins="0-1", want=0),
        test(ins="words and 987", want=0),
        test(ins="0  123", want=0),
        test(ins="  0000000000012345678", want=12345678),
    ]

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
