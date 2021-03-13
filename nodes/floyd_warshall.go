package nodes

// Floyd-Warshall for all shortest paths
// complexity O(V*V*V) + space O(V*V)
// leetcode: 1334
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// setup dp matrix
	infinity := 100000
	dp := make([][]int, n)
	for i := 0; i < n; i++ { // O(n*n)
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j] = 0 // from vertex X to vertex X
				continue
			}
			dp[i][j] = infinity
		}
	}
	for _, edge := range edges { // O(E) = O(n*(n-1)/2)
		from, to, weight := edge[0], edge[1], edge[2]
		if weight > distanceThreshold { // skip too heavy edges
			continue
		}
		dp[from][to] = weight
		dp[to][from] = weight
	}

	// all shortest paths = O(n*n*n)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][k]+dp[k][j] < dp[i][j] {
					dp[i][j] = dp[i][k] + dp[k][j] // go through vertex k
				}
			}
		}
	}

	// no need to check for negative cycles, because weights are all positive

	minCount := n + 1
	minIndex := -1
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if dp[i][j] <= distanceThreshold {
				count++
			}
		}
		if count <= minCount {
			minCount = count
			minIndex = i
		}
	}
	return minIndex
}
