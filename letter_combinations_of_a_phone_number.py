from defs import tests


class Solution:
    def letterCombinations(self, digits: str) -> list[str]:
        if not digits:
            return []

        num_to_letters = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz",
        }

        def backtrack(combination: str, digits: str) -> None:
            if not digits:
                output.append(combination)
            else:
                for l in num_to_letters[digits[0]]:
                    backtrack(combination + l, digits[1:])

        output: list[str] = []
        backtrack("", digits)
        return output

    TESTS = [
        tests(ins="23", want=["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]),
        tests(ins="", want=[]),
        tests(ins="2", want=["a", "b", "c"]),
    ]
