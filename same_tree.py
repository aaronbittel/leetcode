from defs import test, TreeNode
from typing import Optional


class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p and not q:
            return True

        if p and q and p.val == q.val:
            return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)

        return False

    TESTS = [
        test(
            ins=(
                TreeNode(val=1, left=TreeNode(val=2), right=TreeNode(val=3)),
                TreeNode(val=1, left=TreeNode(val=2), right=TreeNode(val=3)),
            ),
            want=True,
        ),
        test(
            ins=(
                TreeNode(val=1, left=TreeNode(val=2)),
                TreeNode(val=1, right=TreeNode(val=2)),
            ),
            want=False,
        ),
        test(
            ins=(
                TreeNode(val=1, left=TreeNode(val=2), right=TreeNode(val=1)),
                TreeNode(val=1, left=TreeNode(val=1), right=TreeNode(val=2)),
            ),
            want=False,
        ),
    ]
