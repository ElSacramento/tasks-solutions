package subsequence

func LongestSubstring(s string, k int) int {
	return longestSubstringRec(s, 0, len(s), k)
}

// O(nlogn), worst O(n2)
func longestSubstringRec(s string, start, end, k int) int {
	if end-start < k {
		return 0
	}

	countInfo := make([]int, 26, 26)
	for i := start; i < end; i++ {
		elem := s[i]
		countInfo[elem-'a']++
	}
	delimiter := -1
	for i := start; i < end; i++ {
		elem := s[i]
		if countInfo[elem-'a'] < k {
			delimiter = i
			break
		}
	}
	if delimiter == -1 {
		return end - start
	}
	return maxInt(longestSubstringRec(s, start, delimiter, k), longestSubstringRec(s, delimiter+1, end, k))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
