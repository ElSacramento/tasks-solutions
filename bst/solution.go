package bst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) insert(val int) {
	if val > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
			return
		}
		n.Right.insert(val)
	} else {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
			return
		}
		n.Left.insert(val)
	}
}

func (n *TreeNode) insertBalanced(val int, heights map[*TreeNode]int) *TreeNode {
	// O(logn) to find place
	if val > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
			heights[n.Right] = 0
			updateHeight(n, heights)
			return n
		}
		n.Right = n.Right.insertBalanced(val, heights)
	} else {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
			heights[n.Left] = 0
			updateHeight(n, heights)
			return n
		}
		n.Left = n.Left.insertBalanced(val, heights)
	}

	updateHeight(n, heights)
	balance := heights[n.Left] - heights[n.Right]
	if balance > 1 {
		// left-heavy
		x := n
		y := n.Left

		// value was added to the right chain of left child
		if val > y.Val {
			// left-right rotation
			xx := y
			yy := y.Right
			leftRotation(xx, yy)
			updateHeight(xx, heights) // child
			y = yy
			x.Left = yy
			rightRotation(x, y)
			updateHeight(x, heights) // child
			updateHeight(y, heights) // parent
			return y
		}
		// right rotation
		rightRotation(x, y)
		updateHeight(x, heights)
		updateHeight(y, heights) // parent
		return y
	} else if balance < -1 {
		// right-heavy
		x := n
		y := n.Right

		// value was added to the left chain of right child
		if val < y.Val {
			// right-left rotation
			xx := y
			yy := y.Left
			rightRotation(xx, yy)
			updateHeight(xx, heights) // child
			y = yy
			x.Right = yy
			leftRotation(x, y)
			updateHeight(x, heights) // child
			updateHeight(y, heights) // parent
			return y
		}
		// left rotation
		leftRotation(x, y)
		updateHeight(x, heights)
		updateHeight(y, heights) // parent
		return y
	}

	return n
}

func (n *TreeNode) traverse(values *[]int) {
	if n.Left != nil {
		n.Left.traverse(values)
	}
	*values = append(*values, n.Val)
	if n.Right != nil {
		n.Right.traverse(values)
	}
}

func CreateBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val > root.Val {
			root.insert(val)
			continue
		}
		root.insert(val)
	}
	return root
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	values := make([]int, 0)
	root.traverse(&values)
	return values
}

func CreateBalancedBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	heights := make(map[*TreeNode]int)
	heights[nil] = -1

	root := &TreeNode{Val: nums[0]}
	heights[root] = 0

	// O(n * logn) + rotations
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		if val > root.Val {
			root = root.insertBalanced(val, heights)
			continue
		}
		root = root.insertBalanced(val, heights)
	}
	return root
}

// AVL tree: balanced bst
// 4 rotations: right, right-left, left, left-right
// for every node: |height(Tree left) - height(Tree right)| <= 1

// return balanced binary search tree
func balanceBST(root *TreeNode) *TreeNode {
	// O(n) + space O(n)
	values := InOrderTraversal(root)

	// space O(n * 2) + O((n * logn) + rotations)
	newRoot := CreateBalancedBST(values)
	return newRoot
}

func leftRotation(x, y *TreeNode) {
	x.Right = y.Left
	y.Left = x
}

func rightRotation(x, y *TreeNode) {
	x.Left = y.Right
	y.Right = x
}

func updateHeight(node *TreeNode, heights map[*TreeNode]int) {
	heights[node] = maxHeight(heights[node.Left], heights[node.Right]) + 1
}

func maxHeight(a, b int) int {
	if a > b {
		return a
	}
	return b
}
