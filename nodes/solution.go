package nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var anc *TreeNode
	pDirection := 0
	qDirection := 0

	next := root
	for next != nil {
		// can find the lowest by value
		anc = next

		if next.Val < p.Val {
			pDirection = 1
		} else if next.Val > p.Val {
			pDirection = -1
		} else {
			pDirection = 0
		}

		if next.Val < q.Val {
			qDirection = 1
		} else if next.Val > q.Val {
			qDirection = -1
		} else {
			qDirection = 0
		}

		directionSum := pDirection + qDirection
		switch directionSum {
		case 0:
			// opposite directions, no more common ancestors
			return anc
		case 1, -1:
			// current node is p or q -> no more common ancestors
			return anc
		case 2:
			next = next.Right
		case -2:
			next = next.Left
		default:
			panic("impossible")
		}
	}

	return anc
}

// better solution
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root.Val < p.Val && root.Val < q.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	if root.Val > p.Val && root.Val > q.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	return root
// }

func traverse(root *TreeNode, val int) {
	prev := root
	current := root
	for current != nil {
		prev = current
		if current.Val > val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	if prev.Val > val {
		prev.Left = &TreeNode{Val: val}
	} else {
		prev.Right = &TreeNode{Val: val}
	}
}

// TODO
func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	for i := 1; i < len(preorder); i++ {
		traverse(root, preorder[i])
	}
	return root
}
