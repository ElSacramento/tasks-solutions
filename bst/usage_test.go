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

func TestMaxLevelSum(t *testing.T) {
	{
		nums := []int{8, 6, 10, 4, 7}
		root := CreateBST(nums)
		require.Equal(t, 2, maxLevelSum(root))
	}
	{
		nums := []int{8, 6, 10, 4, 7, 11}
		root := CreateBST(nums)
		require.Equal(t, 3, maxLevelSum(root))
	}
}

func TestBstFromPreorder(t *testing.T) {
	{
		preorder := []int{4, 6, 5}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{4, 5, 6}, inOrder)
	}
	{
		preorder := []int{1, 3}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 3}, inOrder)
	}
	{
		preorder := []int{3, 2, 1}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 2, 3}, inOrder)
	}
	{
		preorder := []int{2, 1, 3}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 2, 3}, inOrder)
	}
	{
		preorder := []int{8, 5, 1, 7, 10, 12}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 5, 7, 8, 10, 12}, inOrder)
	}
	{
		preorder := []int{8, 5, 1, 6, 7, 10, 9, 12, 11}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 5, 6, 7, 8, 9, 10, 11, 12}, inOrder)
	}
	{
		preorder := []int{8, 5, 3, 1, 6, 10}
		root := bstFromPreorder(preorder)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 3, 5, 6, 8, 10}, inOrder)
	}
}
