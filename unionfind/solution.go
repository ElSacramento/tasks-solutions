package unionfind

type UnionFind struct {
	length   int
	parents  []int       // store indexes of parents for every node
	keyIndex map[int]int // link key to node index
	edges    []PairInt
}

type PairInt struct {
	index1 int
	index2 int
}

func Create() *UnionFind {
	return &UnionFind{
		parents:  make([]int, 0),
		keyIndex: make(map[int]int),
		edges:    make([]PairInt, 0),
	}
}

func calcKey(x, y int) int {
	return x*1000 + y // because x, y can be 300
}

func (uf *UnionFind) AddNode(x, y int) {
	key := calcKey(x, y)

	newIndex := uf.length
	uf.parents = append(uf.parents, newIndex)
	uf.keyIndex[key] = newIndex
	uf.length++
}

func (uf *UnionFind) AddEdge(x1, y1, x2, y2 int) {
	key1 := calcKey(x1, y1)
	key2 := calcKey(x2, y2)

	index1, ok := uf.keyIndex[key1]
	if !ok {
		panic("node not found")
	}
	index2, ok := uf.keyIndex[key2]
	if !ok {
		panic("node not found")
	}
	uf.edges = append(uf.edges, PairInt{index1: index1, index2: index2})
}

func (uf *UnionFind) findParent(index int) int {
	if index >= len(uf.parents) {
		panic("impossible index of node")
	}
	parentInd := uf.parents[index]

	for parentInd != uf.parents[parentInd] {
		parentInd = uf.parents[parentInd]
	}

	uf.parents[index] = parentInd // path compression

	return parentInd
}

func (uf *UnionFind) union(index1, index2 int) {
	uf.parents[index1] = index2
}

func (uf *UnionFind) NumberOfParts() int {
	numberOfParts := uf.length // all nodes are single parts at the start

	for _, edge := range uf.edges {
		parent1 := uf.findParent(edge.index1)
		parent2 := uf.findParent(edge.index2)

		if parent1 == parent2 {
			// almost in one group
			continue
		}
		uf.union(parent1, parent2)
		numberOfParts--
	}
	return numberOfParts
}

// find number of groups, two coordinates are in one group if they are connected by vertical or horizontal line
// used UnionFind for practice
// leetcode: 200
func numIslands(grid [][]byte) int {
	unionFind := Create()

	// O(n * m)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			unionFind.AddNode(i, j)
			if i > 0 && grid[i-1][j] == 1 {
				unionFind.AddEdge(i-1, j, i, j)
			}
			if j > 0 && grid[i][j-1] == 1 {
				unionFind.AddEdge(i, j-1, i, j)
			}
		}
	}

	return unionFind.NumberOfParts()
}
