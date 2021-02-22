package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBST(t *testing.T) {
	{
		nums := []int{8, 6, 10, 4, 7, 3}
		root := CreateBST(nums)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{3, 4, 6, 7, 8, 10}, inOrder)
	}
	{
		nums := []int{1, 2, 3}
		root := CreateBST(nums)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 2, 3}, inOrder)
	}
}

func TestBalanceBST(t *testing.T) {
	{
		nums := []int{8, 6, 10, 4, 7, 3}
		root := CreateBST(nums)
		require.Equal(t, 8, root.Val)
		balanced := balanceBST(root)
		require.Equal(t, 7, balanced.Val)
		inOrder := InOrderTraversal(balanced)
		require.Equal(t, []int{3, 4, 6, 7, 8, 10}, inOrder)
	}
	{
		nums := []int{9, 6, 10, 7, 8}
		root := CreateBST(nums)
		require.Equal(t, 9, root.Val)
		balanced := balanceBST(root)
		require.Equal(t, 7, balanced.Val)
		inOrder := InOrderTraversal(balanced)
		require.Equal(t, []int{6, 7, 8, 9, 10}, inOrder)
	}
	{
		nums := []int{8, 10, 6, 7, 5, 4, 1, 3}
		root := CreateBST(nums)
		require.Equal(t, 8, root.Val)
		balanced := balanceBST(root)
		require.Equal(t, 5, balanced.Val)
		inOrder := InOrderTraversal(balanced)
		require.Equal(t, []int{1, 3, 4, 5, 6, 7, 8, 10}, inOrder)
	}
	{
		nums := []int{1, 15, 14, 17, 7, 2, 3, 12, 9, 11}

		root := CreateBST(nums)
		require.Equal(t, 1, root.Val)
		inOrder := InOrderTraversal(root)
		require.Equal(t, []int{1, 2, 3, 7, 9, 11, 12, 14, 15, 17}, inOrder)

		balanced := balanceBST(root)
		require.Equal(t, 7, balanced.Val)
		inOrder = InOrderTraversal(balanced)
		require.Equal(t, []int{1, 2, 3, 7, 9, 11, 12, 14, 15, 17}, inOrder)
	}
}
