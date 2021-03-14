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

func calcLevel(node *TreeNode, level int, levels map[int]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		return
	}
	nodeSum := 0
	if node.Left != nil {
		nodeSum += node.Left.Val
	}
	if node.Right != nil {
		nodeSum += node.Right.Val
	}
	value, ok := levels[level]
	if !ok {
		levels[level] = nodeSum
	} else {
		levels[level] = value + nodeSum
	}
	if node.Left != nil {
		calcLevel(node.Left, level+1, levels)
	}
	if node.Right != nil {
		calcLevel(node.Right, level+1, levels)
	}
}

// leetcode: 1161
// calculate sum of the elements for every level
// recursion to check every node = O(n)
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelInfo := make(map[int]int)
	levelInfo[1] = root.Val
	calcLevel(root, 2, levelInfo)

	maxSum := levelInfo[1]
	maxLevel := 1
	for k, v := range levelInfo {
		if v > maxSum {
			maxSum = v
			maxLevel = k
		}
	}
	return maxLevel
}
