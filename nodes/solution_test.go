package nodes

import (
	"fmt"
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

func superPrint(head *RandomNode) {
	node := head
	for node != nil {
		fmt.Printf("%p %+v -> ", node, *node)
		node = node.Next
	}
	fmt.Println()
}

func TestCopyRandomList(t *testing.T) {
	{
		tailNode := RandomNode{
			Val:    10,
			Next:   nil,
			Random: nil,
		}
		nodeSecond := RandomNode{
			Val:    20,
			Next:   &tailNode,
			Random: nil,
		}
		nodeFirst := RandomNode{
			Val:    100,
			Next:   &nodeSecond,
			Random: nil,
		}
		rootNode := RandomNode{
			Val:  -1,
			Next: &nodeFirst,
		}
		tailNode.Random = &nodeSecond
		nodeFirst.Random = &tailNode
		rootNode.Random = &rootNode

		newHead := copyRandomList(&rootNode)
		superPrint(&rootNode)
		superPrint(newHead)
	}
}

func getValues(root *ListNode) []int {
	result := make([]int, 0)
	curr := root
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	{
		root1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		}
		root2 := &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		}
		result := addTwoNumbers(root1, root2)
		require.Equal(t, []int{7, 5, 2, 1}, getValues(result))
	}
	{
		root1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		}
		root2 := &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		}
		result := addTwoNumbers(root1, root2)
		require.Equal(t, []int{7, 5, 7}, getValues(result))
	}

}

func TestInorderTraversal(t *testing.T) {
	{
		root := &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Right: nil,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		}
		require.Equal(t, []int{1, 3, 2}, inorderTraversal(root))
	}
	{
		root := &TreeNode{
			Val:   1,
			Right: nil,
			Left: &TreeNode{
				Val:   2,
				Right: nil,
				Left:  nil,
			},
		}
		require.Equal(t, []int{2, 1}, inorderTraversal(root))
	}
	{
		root := &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Right: nil,
				Left:  nil,
			},
		}
		require.Equal(t, []int{1, 2}, inorderTraversal(root))
	}
	{
		leftHand := &TreeNode{
			Val:   2,
			Right: nil,
			Left: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Right: nil,
					Left: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
				},
			},
		}
		rightHand := &TreeNode{
			Val:   5,
			Right: nil,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}
		root := &TreeNode{
			Val:   1,
			Right: rightHand,
			Left:  leftHand,
		}
		require.Equal(t, []int{4, 8, 7, 2, 1, 3, 5}, inorderTraversal(root))
	}
	{
		root := &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}
		require.Equal(t, []int{1}, inorderTraversal(root))
	}
}

// func toArray(root *TreeNode, values *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	toArray(root.Left, values)
// 	values = append(values, root.Val)
// 	toArray(root.Right, values)
// }
//
// func TestBstFromPreorder(t *testing.T) {
// 	t.Run("tree number one", func(t *testing.T) {
// 		root := bstFromPreorder([]int{4, 3, 1, 8, 7})
// 		result := make([]int, 0)
// 		toArray(root, result)
// 		require.Equal(t, []int{4, 3, 1}, result)
// 	})
// }

func TestDeleteNode(t *testing.T) {
	{
		head := &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: -2,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		}
		deleteNode(head.Next)
		require.Equal(t, []int{10, 7}, getValues(head))
	}
	{
		head := &ListNode{
			Val: 10,
			Next: &ListNode{
				Val:  -2,
				Next: nil,
			},
		}
		deleteNode(head)
		require.Equal(t, []int{-2}, getValues(head))
	}
	{
		head := &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: -2,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
		}
		deleteNode(head.Next.Next)
		require.Equal(t, []int{10, -2, 0}, getValues(head))
	}
}

func TestMaxDepth(t *testing.T) {
	{
		root := &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
			},
		}
		require.Equal(t, 2, maxDepth(root))
	}
	{
		root := &TreeNode{
			Val:   1,
			Right: nil,
			Left: &TreeNode{
				Val: 2,
			},
		}
		require.Equal(t, 2, maxDepth(root))
	}
	{
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		require.Equal(t, 2, maxDepth(root))
	}
	{
		leftHand := &TreeNode{
			Val:   2,
			Right: nil,
			Left: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Right: nil,
					Left: &TreeNode{
						Val: 8,
					},
				},
			},
		}
		rightHand := &TreeNode{
			Val:   5,
			Right: nil,
			Left: &TreeNode{
				Val: 3,
			},
		}
		root := &TreeNode{
			Val:   1,
			Right: rightHand,
			Left:  leftHand,
		}
		require.Equal(t, 5, maxDepth(root))
	}
	{
		root := &TreeNode{
			Val: 1,
		}
		require.Equal(t, 1, maxDepth(root))
	}
	{
		require.Equal(t, 0, maxDepth(nil))
	}
}
