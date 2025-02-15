# Solution:
# 2 pointers
# time-complexity: O(n**2)
# space-complexity: O(n)

from defs import tests


class Solution:
    TESTS = [
        tests(ins="babad", want="bab"),
        tests(ins="cbbd", want="bb"),
        tests(ins="a", want="a"),
    ]

    def longestPalindrome(self, s: str) -> str:
        longest = s[0]
        for i in range(len(s)):
            # if remaining length is less than longest found palindrom; return
            if len(s) - i + 1 < len(longest):
                return longest
            for j in range(len(s) - 1, i, -1):
                # early break if length of found palindrom > current window
                # if j - i + 1 < len(longest):
                #     break
                if s[j] != s[i]:
                    continue
                if is_palindrom(s[i : j + 1]):
                    if j - i + 1 > len(longest):
                        longest = s[i : j + 1]
                    break
        return longest


def is_palindrom(s: str) -> bool:
    half = len(s) // 2
    if len(s) % 2 == 0:
        return s[:half] == s[: half - 1 : -1]
    return s[:half] == s[:half:-1]


# def is_palindrom(s: str) -> bool:
#     l = len(s)
#     for i in range(l // 2):
#         if s[i] != s[l - 1 - i]:
#             return False
#     return True
