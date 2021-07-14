package nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n) - worst, for balanced O(logn)
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
// O(logn)
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode: 2
// O(n)
// sum two linked lists items
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result     ListNode
		tempo, sum int
	)
	pointer := &result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if tempo != 0 {
			sum += tempo
			tempo = 0
		}
		if sum >= 10 {
			sum %= 10
			tempo = 1
		}
		pointer.Val = sum
		if l1 != nil || l2 != nil {
			pointer.Next = &ListNode{}
			pointer = pointer.Next
		}
		sum = 0
	}
	if tempo != 0 {
		pointer.Next = &ListNode{Val: tempo}
	}
	return &result
}

type stack []*TreeNode

func (s *stack) add(n *TreeNode) {
	old := *s
	old = append(old, n)
	*s = old
}

func (s *stack) pop() *TreeNode {
	old := *s
	n := len(old)
	value := old[n-1]
	old = old[:n-1]
	*s = old
	return value
}

// leetcode: 94
// binary tree, no rules
// O(n) + space for stack O(n) + map O(n)
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nodesStack := make(stack, 0)
	nodesStack.add(root)

	visited := make(map[*TreeNode]struct{})
	result := make([]int, 0)
	current := root
	// O(n)
	for current != nil {
		if _, ok := visited[current]; ok {
			if len(nodesStack) == 0 {
				current = nil
			} else {
				current = nodesStack.pop()
			}
			continue
		}

		if current.Left != nil {
			if _, ok := visited[current.Left]; !ok {
				nodesStack.add(current)
				current = current.Left
				continue
			}
		}

		// inOrder
		result = append(result, current.Val)
		visited[current] = struct{}{}

		if current.Right != nil {
			if _, ok := visited[current.Right]; !ok {
				current = current.Right
				continue
			}
		}

		if len(nodesStack) == 0 {
			current = nil
		} else {
			current = nodesStack.pop()
		}
	}
	return result
}

// leetcode: 237
// without head ref
func deleteNode(node *ListNode) {
	*node = *node.Next
}

func findDepth(node *TreeNode, depth int) int {
	leftDepth := depth
	rightDepth := depth
	if node.Left != nil {
		leftDepth = findDepth(node.Left, depth+1)
	}
	if node.Right != nil {
		rightDepth = findDepth(node.Right, depth+1)
	}
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

// leetcode: 104
// visit all nodes in recursive way: O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 1)
}

func traverse(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}
	if len(*result) <= level {
		tempo := *result
		tempo = append(tempo, []int{node.Val})
		*result = tempo
	} else {
		tempo := *result
		tempo[level] = append(tempo[level], node.Val)
		*result = tempo
	}
	traverse(node.Left, level+1, result)
	traverse(node.Right, level+1, result)
}

// leetcode: 102
// visit all nodes O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	result = append(result, []int{root.Val})
	traverse(root.Left, 1, &result)
	traverse(root.Right, 1, &result)
	return result
}

// leetcode: 206
// reverse list with recursion = O(n) just visit all nodes
// approach with iteration reduce memory usage
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	reversed := reverse(head)
	head.Next = nil
	return reversed
}

func reverse(node *ListNode) *ListNode {
	next := node.Next
	if next != nil {
		newHead := reverse(next)
		next.Next = node
		return newHead
	}
	return node
}
