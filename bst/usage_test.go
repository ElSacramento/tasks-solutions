package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertIntoBST(t *testing.T) {
	{
		nums := []int{8, 6, 10, 4, 7, 3}
		root := CreateBST(nums)
		require.Equal(t, 8, root.Val)
		updated := insertIntoBST(root, 9)
		inOrder := InOrderTraversal(updated)
		require.Equal(t, []int{3, 4, 6, 7, 8, 9, 10}, inOrder)

		updated = insertIntoBST(root, 5)
		inOrder = InOrderTraversal(updated)
		require.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 10}, inOrder)

		updated = insertIntoBST(root, 20)
		inOrder = InOrderTraversal(updated)
		require.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 10, 20}, inOrder)
	}
	{
		nums := []int{8}
		root := CreateBST(nums)
		require.Equal(t, 8, root.Val)
		updated := insertIntoBST(root, 9)
		inOrder := InOrderTraversal(updated)
		require.Equal(t, []int{8, 9}, inOrder)
	}
}
