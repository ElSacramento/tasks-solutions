package nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLowestCommonAncestor(t *testing.T) {
	t.Run("root is p", func(t *testing.T) {
		var (
			leaf = &TreeNode{Val: 1}
			root = &TreeNode{Val: 2, Left: leaf}
		)
		result := lowestCommonAncestor(root, root, leaf)
		require.Equal(t, 2, result.Val)
	})
	t.Run("root is q", func(t *testing.T) {
		var (
			leaf = &TreeNode{Val: 3}
			root = &TreeNode{Val: 2, Left: leaf}
		)
		result := lowestCommonAncestor(root, leaf, root)
		require.Equal(t, 2, result.Val)
	})
	t.Run("medium tree", func(t *testing.T) {
		var (
			thirdLevelLeafLeft3   = &TreeNode{Val: 8}
			secondLevelLeafLeft2  = &TreeNode{Val: 9, Left: thirdLevelLeafLeft3}
			secondLevelLeafRight2 = &TreeNode{Val: 12}
			secondLevelLeafLeft   = &TreeNode{Val: 2}
			secondLevelLeafRight  = &TreeNode{Val: 5}
			firstLevelLeafLeft    = &TreeNode{Val: 4, Left: secondLevelLeafLeft, Right: secondLevelLeafRight}
			firstLevelLeafRight   = &TreeNode{Val: 10, Left: secondLevelLeafLeft2, Right: secondLevelLeafRight2}
			root                  = &TreeNode{Val: 6, Left: firstLevelLeafLeft, Right: firstLevelLeafRight}
		)
		result := lowestCommonAncestor(root, thirdLevelLeafLeft3, secondLevelLeafRight2)
		require.Equal(t, 10, result.Val)

		result = lowestCommonAncestor(root, firstLevelLeafLeft, firstLevelLeafRight)
		require.Equal(t, 6, result.Val)

		result = lowestCommonAncestor(root, secondLevelLeafLeft, secondLevelLeafRight)
		require.Equal(t, 4, result.Val)

		result = lowestCommonAncestor(root, firstLevelLeafRight, secondLevelLeafRight2)
		require.Equal(t, 10, result.Val)
	})
}

func toArray(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	toArray(root.Left, values)
	values = append(values, root.Val)
	toArray(root.Right, values)
}

func TestBstFromPreorder(t *testing.T) {
	t.Run("tree number one", func(t *testing.T) {
		root := bstFromPreorder([]int{4, 3, 1, 8, 7})
		result := make([]int, 0)
		toArray(root, result)
		require.Equal(t, []int{4, 3, 1}, result)
	})
}
