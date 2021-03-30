package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxAverageSubtree(t *testing.T) {
	{
		nodeLeaf1 := Node{value: 11}
		nodeLeaf2 := Node{value: 3}
		nodeLeaf3 := Node{value: 1}
		nodeSubtree := Node{value: 5, children: []*Node{&nodeLeaf2, &nodeLeaf3}}
		root := Node{value: 12, children: []*Node{&nodeLeaf1, &nodeSubtree}}
		nodeRoot, sum := maxAverageSubtree(&root)
		require.Equal(t, 12, nodeRoot)
		require.Equal(t, 6, sum)
	}
	{
		nodeSubtree := Node{value: 20, children: []*Node{{value: 10}, {value: 14}}}
		root := Node{value: 7, children: []*Node{{value: 5}, &nodeSubtree}}
		nodeRoot, sum := maxAverageSubtree(&root)
		require.Equal(t, 20, nodeRoot)
		require.Equal(t, 14, sum)
	}
}
