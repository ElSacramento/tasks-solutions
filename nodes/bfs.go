package nodes

type pairInt struct {
	x int
	y int
}

func possibleNeighbours(x, y, n int) []pairInt {
	result := make([]pairInt, 0)

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i >= 0 && i < n && j >= 0 && j < n {
				result = append(result, pairInt{x: i, y: j})
			}
		}
	}
	return result
}

// breadth-first-search
// leetcode: 1091
// go through cells where value is 0
func shortestPathBinaryMatrix(grid [][]int) int {
	last := len(grid) - 1
	if grid[0][0] == 1 || grid[last][last] == 1 {
		return -1
	}

	// number of zeros = V -> n*n at max
	queue := make([]pairInt, 0)        // space O(V)
	distances := make(map[pairInt]int) // space O(V)

	headIndex := 0
	start := pairInt{x: 0, y: 0}
	queue = append(queue, start)
	distances[start] = 1

	// O(V + E)
	for headIndex < len(queue) {
		node := queue[headIndex]
		headIndex++

		x, y := node.x, node.y

		// up, down, left, right
		neighbours := possibleNeighbours(x, y, len(grid))
		for _, nextNode := range neighbours {
			if grid[nextNode.x][nextNode.y] == 0 {
				// already added in queue
				if v, ok := distances[nextNode]; ok {
					if distances[node]+1 < v {
						distances[nextNode] = distances[node] + 1
					}
					continue
				}
				distances[nextNode] = distances[node] + 1
				queue = append(queue, nextNode)
			}
		}
	}

	end := pairInt{x: last, y: last}
	if v, ok := distances[end]; ok {
		return v
	}
	return -1
}
