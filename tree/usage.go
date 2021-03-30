package tree

import (
	"fmt"
)

type Node struct {
	value    int
	children []*Node
}

type PairInt struct {
	sum   int
	count int
}

func goDeep(root, node *Node, subtrees map[*Node]PairInt) {
	fmt.Printf("node: %v\n", node.value)

	if len(node.children) == 0 {
		subtrees[root] = PairInt{sum: subtrees[root].sum + node.value, count: subtrees[root].count + 1}
		fmt.Printf("add leaf: %+v\n", subtrees[root])
		return
	}

	subtrees[node] = PairInt{sum: node.value, count: 1}
	fmt.Printf("subtree: %+v\n", subtrees[node])

	for _, ch := range node.children {
		goDeep(node, ch, subtrees)
	}

	// use result from children
	if root != node {
		subtrees[root] = PairInt{sum: subtrees[root].sum + subtrees[node].sum, count: subtrees[root].count + subtrees[node].count}
		fmt.Printf("updated prevnode: %+v\n", subtrees[root])
	}
}

// subtree with max average sum
// can be more than 2 children
// DP: calculate for leaf and propose higher with recursion calls
// O(n) need to visit all nodes
func maxAverageSubtree(root *Node) (int, int) {
	subtrees := make(map[*Node]PairInt)
	goDeep(root, root, subtrees)

	max := 0
	var node *Node
	for n, v := range subtrees {
		if v.sum/v.count > max {
			node = n
			max = v.sum / v.count
		}
	}

	if node == nil {
		panic("impossible")
	}

	return node.value, max
}
