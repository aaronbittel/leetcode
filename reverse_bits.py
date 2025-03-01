from defs import test


class Solution:
    def reverseBits(self, n: int) -> int:
        result = 0b0

        for _ in range(32):
            bit = n & 1
            result = (result << 1) | bit
            n >>= 1

        return result

    TESTS = [
        test(ins=0b00000010100101000001111010011100, want=964176192),
        test(ins=0b11111111111111111111111111111101, want=3221225471),
    ]
