package nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n)
// leetcode: 235
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

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// leetcode: 138
// O(n) + space for new list and tracking random links
func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	index := 0
	randomNodes := make(map[*RandomNode][]int, 1000)
	newNodes := make(map[int]*RandomNode, 1000)

	newHead := RandomNode{
		Val:  head.Val,
		Next: &RandomNode{},
	}
	if head.Next == nil {
		newHead.Next = nil
		if head.Random != nil {
			newHead.Random = &newHead
		}
		return &newHead
	}

	if head.Random != nil {
		// to track random links
		randomNodes[head.Random] = append(randomNodes[head.Random], index)
	}
	newNodes[index] = &newHead

	index++
	newCurrentNode := newHead.Next
	currentNode := head.Next
	// O(n)
	for currentNode != nil {
		value := currentNode.Val
		randomLink := currentNode.Random
		if randomLink != nil {
			randomNodes[randomLink] = append(randomNodes[randomLink], index)
		}
		newCurrentNode.Val = value
		newNodes[index] = newCurrentNode

		currentNode = currentNode.Next
		if currentNode != nil {
			newCurrentNode.Next = &RandomNode{}
			newCurrentNode = newCurrentNode.Next
		}
		index++
	}

	index = 0
	currentNode = head
	// O(n)
	for currentNode != nil {
		// update random nodes
		if nodes, ok := randomNodes[currentNode]; ok {
			for _, n := range nodes {
				newNodes[n].Random = newNodes[index]
			}
			delete(randomNodes, currentNode)
		}

		currentNode = currentNode.Next
		index++
	}
	return &newHead
}

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
