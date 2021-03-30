package other

// leetcode: 1328
// change minimum chars to break a palindrome + lexicographically smaller
// O(n)
func breakPalindrome(palindrome string) string {
	bytesSt := []byte(palindrome)

	if len(bytesSt) == 1 {
		return ""
	}

	maxLeft := -1
	for i, el := range bytesSt {
		// change middle element is useless
		if len(bytesSt)%2 == 1 && i == len(bytesSt)/2 {
			continue
		}
		// change first elem, that is not 'a' to 'a'
		if el > 'a' {
			maxLeft = i
			break
		}
	}
	// aaaa -> aaab
	if maxLeft == -1 {
		bytesSt[len(bytesSt)-1] = 'b'
		return string(bytesSt)
	}

	bytesSt[maxLeft] = 'a'
	return string(bytesSt)
}

// leetcode: 3
// sliding window O(N)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	maxCount := 1
	indexes := make(map[byte]int, 0) // space O(n)
	indexes[s[0]] = 1                // index + 1
	for right != len(s)-1 {
		right += 1

		index, ok := indexes[s[right]]
		if !ok || index == 0 {
			indexes[s[right]] = right + 1
			continue
		}

		if right-left > maxCount {
			maxCount = right - left
		}
		for left != index {
			indexes[s[left]] = 0
			left += 1
		}
		indexes[s[right]] = right + 1
	}
	if right-left+1 > maxCount {
		maxCount = right - left + 1
	}
	return maxCount
}
