package bst

// leetcode: 701
// O(logn) - find leaf to insert value
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	prev := root
	curr := root
	for curr != nil {
		prev = curr
		if val > curr.Val {
			curr = curr.Right
			continue
		}
		curr = curr.Left
	}

	if val > prev.Val {
		prev.Right = &TreeNode{Val: val}
	} else {
		prev.Left = &TreeNode{Val: val}
	}

	return root
}
