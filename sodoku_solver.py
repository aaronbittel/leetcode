from defs import test


class Solution:
    def solveSudoku(self, board: list[list[str]]) -> None:
        def solve(board: list[list[str]]) -> bool:
            for row in range(9):
                for col in range(9):
                    if board[row][col] != ".":
                        continue

                    for n in filter(
                        lambda x: is_valid(board, row, col, x), "123456789"
                    ):
                        board[row][col] = n
                        if solve(board):
                            return True
                        board[row][col] = "."

                    return False

            return True

        solve(board)

    TESTS = [
        test(
            ins=[
                ["5", "3", ".", ".", "7", ".", ".", ".", "."],
                ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                [".", "9", "8", ".", ".", ".", ".", "6", "."],
                ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                [".", "6", ".", ".", ".", ".", "2", "8", "."],
                [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                [".", ".", ".", ".", "8", ".", ".", "7", "9"],
            ],
            want=[
                ["5", "3", "4", "6", "7", "8", "9", "1", "2"],
                ["6", "7", "2", "1", "9", "5", "3", "4", "8"],
                ["1", "9", "8", "3", "4", "2", "5", "6", "7"],
                ["8", "5", "9", "7", "6", "1", "4", "2", "3"],
                ["4", "2", "6", "8", "5", "3", "7", "9", "1"],
                ["7", "1", "3", "9", "2", "4", "8", "5", "6"],
                ["9", "6", "1", "5", "3", "7", "2", "8", "4"],
                ["2", "8", "7", "4", "1", "9", "6", "3", "5"],
                ["3", "4", "5", "2", "8", "6", "1", "7", "9"],
            ],
        ),
    ]


def is_valid(board: list[list[str]], row: int, col: int, num: str) -> bool:
    if num in board[row]:
        return False

    if num in [board[r][col] for r in range(9)]:
        return False

    box_y, box_x = 3 * (row // 3), 3 * (col // 3)
    for y in range(box_y, box_y + 3):
        for x in range(box_x, box_x + 3):
            if num == board[y][x]:
                return False

    return True


def print_board(board: list[list[str]]):
    for row in range(9):
        print(" ".join(board[row]))
    print()


board = Solution.TESTS[0].ins

print_board(board)
Solution().solveSudoku(board)
print_board(board)
