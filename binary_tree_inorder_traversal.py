from typing import Optional

from defs import TreeNode, test


class Solution:
    # Left-Root-Right
    def inorderTraversal(self, root: Optional[TreeNode]) -> list[int]:
        if root is None:
            return []

        res: list[int] = []

        def inorder(head: TreeNode):
            if head.left is not None:
                inorder(head.left)
            res.append(head.val)
            if head.right is not None:
                inorder(head.right)

        inorder(root)
        return res

    TESTS = [
        test(
            ins=TreeNode(val=1, right=TreeNode(val=2, left=TreeNode(3))),
            want=[1, 3, 2],
        ),
        test(
            ins=TreeNode(val=1, left=TreeNode(val=2), right=TreeNode(3)),
            want=[2, 1, 3],
        ),
        test(
            ins=TreeNode(
                val=1,
                left=TreeNode(
                    val=2,
                    left=TreeNode(4),
                    right=TreeNode(val=5, left=TreeNode(6), right=TreeNode(7)),
                ),
                right=TreeNode(val=3, right=TreeNode(val=8, left=TreeNode(9))),
            ),
            want=[4, 2, 6, 5, 7, 1, 3, 9, 8],
        ),
    ]
