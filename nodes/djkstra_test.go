package nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberOfShortestPaths(t *testing.T) {
	{
		graph := []*Node{
			{neighbours: []Edge{
				{index: 1, weight: 10},
				{index: 3, weight: 4},
				{index: 7, weight: 10},
			}},
			{neighbours: []Edge{
				{index: 2, weight: 2},
				{index: 3, weight: 8},
				{index: 0, weight: 10},
			}},
			{neighbours: []Edge{
				{index: 1, weight: 2},
				{index: 4, weight: 10},
				{index: 5, weight: 4},
			}},
			{neighbours: []Edge{
				{index: 0, weight: 4},
				{index: 1, weight: 8},
				{index: 4, weight: 4},
				{index: 8, weight: 20},
			}},
			{neighbours: []Edge{
				{index: 2, weight: 10},
				{index: 3, weight: 4},
				{index: 6, weight: 2},
			}},
			{neighbours: []Edge{
				{index: 2, weight: 4},
				{index: 6, weight: 8},
			}},
			{neighbours: []Edge{
				{index: 4, weight: 2},
				{index: 5, weight: 8},
			}},
			{neighbours: []Edge{
				{index: 8, weight: 13},
				{index: 0, weight: 10},
			}},
			{neighbours: []Edge{
				{index: 7, weight: 13},
				{index: 3, weight: 20},
			}},
		}
		require.Equal(t, 3, NumberOfShortestPaths(graph, 1, 6))
		require.Equal(t, 2, NumberOfShortestPaths(graph, 1, 4))
	}
}

func TestNetworkDelayTime(t *testing.T) {
	{
		edges := [][]int{
			{2, 1, 1},
			{2, 3, 1},
			{3, 4, 1},
		}
		require.Equal(t, 2, networkDelayTime(edges, 4, 2))
	}
	{
		edges := [][]int{
			{1, 2, 1},
		}
		require.Equal(t, 1, networkDelayTime(edges, 2, 1))
		require.Equal(t, -1, networkDelayTime(edges, 2, 2))
	}
	{
		edges := [][]int{
			{1, 2, 1},
			{3, 4, 1},
		}
		require.Equal(t, -1, networkDelayTime(edges, 4, 1))
	}
	{
		edges := [][]int{
			{2, 1, 1},
			{2, 3, 2},
			{2, 4, 3},
			{2, 5, 2},
			{1, 6, 1},
			{5, 6, 1},
		}
		require.Equal(t, 3, networkDelayTime(edges, 6, 2))
	}
}

func TestMinimumEffortPath(t *testing.T) {
	{
		heights := [][]int{
			{1, 2, 2},
			{3, 8, 2},
			{5, 3, 5},
		}
		require.Equal(t, 2, minimumEffortPath(heights))
	}
	{
		heights := [][]int{
			{1, 2, 3},
			{3, 8, 4},
			{5, 3, 5},
		}
		require.Equal(t, 1, minimumEffortPath(heights))
	}
	{
		heights := [][]int{
			{1, 2, 1},
			{1, 2, 2},
			{3, 1, 1},
		}
		require.Equal(t, 1, minimumEffortPath(heights))
	}
	{
		heights := [][]int{
			{1, 2, 1, 1, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1},
			{1, 1, 1, 2, 1},
		}
		require.Equal(t, 0, minimumEffortPath(heights))
	}
	{
		heights := [][]int{
			{1, 10, 6, 7, 9, 10, 4, 9},
		}
		require.Equal(t, 9, minimumEffortPath(heights))
	}
	{
		heights := [][]int{
			{1},
			{2},
			{8},
			{4},
			{8},
		}
		require.Equal(t, 6, minimumEffortPath(heights))
	}
}

func TestMaxProbability(t *testing.T) {
	{
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		prob := []float64{0.5, 0.5, 0.2}
		require.Equal(t, 0.25, maxProbability(3, edges, prob, 0, 2))
	}
	{
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		prob := []float64{0.5, 0.5, 0.3}
		require.Equal(t, 0.3, maxProbability(3, edges, prob, 0, 2))
	}
	{
		edges := [][]int{{0, 1}}
		prob := []float64{0.5}
		require.Equal(t, float64(0), maxProbability(3, edges, prob, 0, 2))
	}
	{
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {0, 5}, {5, 4}}
		prob := []float64{0.1, 0.2, 0.3, 0.4, 0.005, 0.07}
		require.Equal(t, 0.0024000000000000002, maxProbability(6, edges, prob, 0, 4))
	}
	{
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}
		prob := []float64{0.1, 0.7, 0.1, 0.8}
		require.Equal(t, 0.08000000000000002, maxProbability(4, edges, prob, 0, 3))
	}
	{
		edges := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 2}}
		prob := []float64{0.3, 0.5, 0.4, 0.4}
		require.Equal(t, 0.16000000000000003, maxProbability(4, edges, prob, 0, 2))
	}
}
