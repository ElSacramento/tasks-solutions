package subsequence

// leetcode: 395
// find longest substring, where every character is repeated at least k times
func LongestSubstring(s string, k int) int {
	return longestSubstringRec(s, 0, len(s), k)
}

// O(nlogn), worst O(n2)
func longestSubstringRec(s string, start, end, k int) int {
	if end-start < k {
		return 0
	}

	countInfo := make([]int, 26)
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

// leetcode: 14
// O(n*n)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := 201
	var checkSt string
	// O(n)
	for _, st := range strs {
		if len(st) < minLen {
			checkSt = st
			minLen = len(st)
		}
	}

	if minLen == 0 {
		return ""
	}

	// check first element before
	{
		checkElem := checkSt[0]
		wrong := false

		// O(n)
		for _, st := range strs {
			if st[0] != checkElem {
				wrong = true
				break
			}
		}

		if wrong {
			return ""
		}
	}

	start := 1
	end := false
	// O(n*k)
	for !end && start < len(checkSt) {
		checkElem := checkSt[start]
		wrong := false

		for _, st := range strs {
			if st[start] != checkElem {
				wrong = true
				break
			}
		}

		if wrong {
			end = true
			continue
		}
		start += 1
	}

	prefix := checkSt[:start]
	return prefix
}
