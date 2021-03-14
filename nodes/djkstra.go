package nodes

import (
	"container/heap"
)

type Node struct {
	neighbours []Edge
}

type Edge struct {
	index  int
	weight int
}

type Item struct {
	nodeIndex int
	priority  int // first-priority is higher than second-priority - minHeap
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*p = old[0 : n-1]
	return item
}

// https://www.baeldung.com/cs/graph-number-of-shortest-paths
// weighted graph with undirected edges
func NumberOfShortestPaths(nodes []*Node, startNodeIndex, endNodeIndex int) int {
	distances := make([]int, len(nodes)) // space O(V)
	paths := make([]int, len(nodes))     // space O(V)

	paths[startNodeIndex] = 1     // from start to start
	distances[startNodeIndex] = 0 // from start to start

	priorityQueue := make(PriorityQueue, 0) // space O(V)
	heap.Init(&priorityQueue)
	heap.Push(&priorityQueue, &Item{nodeIndex: startNodeIndex, priority: 0}) // O(logV)
	// priority is the distance

	// O(E + VlogV) - visit all edges + put every vertex in the priority queue
	for priorityQueue.Len() != 0 {
		queueItem := heap.Pop(&priorityQueue).(*Item)
		currentNodeIndex := queueItem.nodeIndex
		currentNodeDistance := queueItem.priority

		node := nodes[currentNodeIndex]
		for _, edge := range node.neighbours {
			neighbourIndex := edge.index
			// haven't visited
			if distances[neighbourIndex] == 0 && neighbourIndex != startNodeIndex {
				heap.Push(&priorityQueue, &Item{nodeIndex: neighbourIndex, priority: currentNodeDistance + edge.weight})
				paths[neighbourIndex] = paths[currentNodeIndex]
				distances[neighbourIndex] = currentNodeDistance + edge.weight
				continue
			}
			// for already visited
			if distances[neighbourIndex] == currentNodeDistance+edge.weight {
				paths[neighbourIndex] += paths[currentNodeIndex]
				continue
			}
			// current distance to the neighbour is greater
			if distances[neighbourIndex] > currentNodeDistance+edge.weight {
				heap.Push(&priorityQueue, &Item{nodeIndex: neighbourIndex, priority: currentNodeDistance + edge.weight})
				paths[neighbourIndex] = paths[currentNodeIndex]
				distances[neighbourIndex] = currentNodeDistance + edge.weight
			}
		}
	}
	return paths[endNodeIndex]
}

// maximum time to visit all nodes
// leetcode: 743
func networkDelayTime(times [][]int, n int, k int) int {
	neighbours := make(map[int][]int, n)
	for i, v := range times { // O(E)
		vertex := v[0]
		neighbours[vertex] = append(neighbours[vertex], i)
	}
	// 2 -> [0, 1], 3 -> [2]

	// corner case - no edges from node[k]
	if edges, ok := neighbours[k]; !ok || len(edges) == 0 {
		return -1
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{priority: 0, nodeIndex: k})

	distances := make(map[int]int)

	for pq.Len() != 0 {
		item := heap.Pop(&pq).(*Item)
		distance := item.priority

		// already visited with less or equal time
		if value, ok := distances[item.nodeIndex]; ok && value <= distance {
			continue
		}
		distances[item.nodeIndex] = distance

		if nbors, ok := neighbours[item.nodeIndex]; ok {
			for _, edgeIndex := range nbors { // O(E*logV)
				edge := times[edgeIndex]
				to := edge[1]
				// already visited
				if _, ok := distances[to]; ok {
					continue
				}
				cost := edge[2]

				heap.Push(&pq, &Item{priority: distance + cost, nodeIndex: to}) // O(logV)
			}
		}
	}

	if len(distances) != n {
		return -1
	}

	maxTime := -1
	for _, d := range distances { // O(V)
		if d > maxTime {
			maxTime = d
		}
	}
	return maxTime
}

type DistancePriorityQueue [][3]int

func (d DistancePriorityQueue) Len() int {
	return len(d)
}

func (d DistancePriorityQueue) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d DistancePriorityQueue) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d *DistancePriorityQueue) Push(x interface{}) {
	item := x.([3]int)
	*d = append(*d, item)
}

func (d *DistancePriorityQueue) Pop() interface{} {
	old := *d
	n := len(old)
	item := old[n-1]
	*d = old[0 : n-1]
	return item
}

// leetcode: 1631
// find path with minimum diff between heights
func minimumEffortPath(heights [][]int) int {
	queue := make(DistancePriorityQueue, 0) // space O(n*n)
	distances := make(map[[2]int]int)       // space(n*n)
	rowsLen := len(heights)
	columnsLen := len(heights[0])

	heap.Push(&queue, [3]int{0, 0, 0})
	for queue.Len() != 0 { // O(n*n)
		item := heap.Pop(&queue).([3]int)
		d, row, column := item[0], item[1], item[2]

		if v, ok := distances[[2]int{row, column}]; ok && v <= d { // already visited with min distance
			continue
		}

		distances[[2]int{row, column}] = d
		if row == rowsLen-1 && column == columnsLen-1 {
			break
		}

		// up, down, left, right
		pairs := possiblePairs(row, column)
		for _, pair := range pairs {
			i, j := pair[0], pair[1]
			if i < 0 || j < 0 || i >= rowsLen || j >= columnsLen {
				continue
			}
			if _, ok := distances[[2]int{i, j}]; ok { // already visited
				continue
			}

			dist := abs(heights[i][j] - heights[row][column])
			priority := max(d, dist)
			heap.Push(&queue, [3]int{priority, i, j})
		}
	}
	return distances[[2]int{rowsLen - 1, columnsLen - 1}]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func possiblePairs(row, column int) [][2]int {
	return [][2]int{
		{row, column - 1},
		{row, column + 1},
		{row - 1, column},
		{row + 1, column},
	}
}
