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
