package nodes

func dfs(i int, M [][]int, seen []bool) {
	for j := 0; j < len(M[i]); j++ {
		if M[i][j] == 1 && !seen[j] {
			seen[j] = true
			dfs(j, M, seen)
		}
	}
}

// leetcode: 547
// dfs with matrix graph
func findCircleNum(M [][]int) int {
	seen := make([]bool, len(M))
	count := 0

	for i := 0; i < len(M); i++ {
		if !seen[i] {
			count++
			dfs(i, M, seen)
		}
	}
	return count
}
