import logging
from functools import cache
from defs import tests

logging.basicConfig(format="", level=logging.INFO)


class Solution:
    TESTS = [
        tests(ins=("aa", "a"), out=False),
        tests(ins=("aa", "a*"), out=True),
        tests(ins=("ab", ".*"), out=True),
        tests(ins=("aab", "c*a*b"), out=True),
        tests(ins=("mississippi", "mis*is*ip*."), out=True),
        tests(ins=("ppi", "."), out=False),
        tests(ins=("ab", ".*c"), out=False),
        tests(ins=("aaa", "a.a"), out=True),
        tests(ins=("a", "ab*"), out=True),
        tests(ins=("a", ".*..a*"), out=False),
        tests(ins=("a", "ab*"), out=True),
        tests(ins=("abcdede", "ab.*de"), out=True),
        tests(ins=("aaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*"), out=False),
    ]

    @cache
    def isMatch(self, s: str, p: str) -> bool:
        if s == p:
            return True
        logging.debug(f"calling: {s = }, {p = }")
        s_i, p_i = 0, 0
        while s_i < len(s) and p_i < len(p) - 1:
            ch = s[s_i] if s_i < len(s) else ""

            choices: list[str] = []
            logging.debug(f"matching: {ch = } and ({p[p_i] = }, {p[p_i+1] = })")
            match (p[p_i], p[p_i + 1]):
                case (".", "*"):
                    logging.debug(f"{s = }, {s_i = }, {s[s_i:] = }")
                    for i in range(len(s) - s_i + 1):
                        choices.append(s[s_i : s_i + i])
                case (prev, "*"):
                    choices.append("")
                    k = 0
                    while s_i + k < len(s) and s[s_i + k] == prev:
                        choices.append((k + 1) * prev)
                        k += 1
                case (".", _):
                    logging.debug("found .")
                case (cur, _):
                    logging.debug(f"\t{cur = } ==? {ch = }")
                    if ch != cur:
                        logging.debug("\tNOT EQUAL -> returning false")
                        return False

            if choices:
                logging.debug(f"{choices = }")
                return any(
                    self.isMatch(s[s_i:], choice + p[p_i + 2 :]) for choice in choices
                )

            s_i += 1
            p_i += 1

        pattern_left = p[p_i:]
        str_left = s[s_i:]
        logging.debug(f"remaining str: {str_left = }, pattern: {pattern_left = }")

        if str_left == pattern_left:
            return True

        if len(str_left) > 0 and len(pattern_left) == 0:
            return False

        if str_left == "":
            if len(pattern_left) % 2 == 1:
                return False
            return all(c == "*" for c in pattern_left[1::2])

        if len(str_left) == 1 and pattern_left == ".":
            return True

        return False
