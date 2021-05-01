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

// leetcode: 17
// can be solved with dfs or recursion
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	letters := map[rune][]string{
		'0': {},
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	if len(digits) == 1 {
		return letters[rune(digits[0])]
	}

	numbers := make([]rune, 4)
	for i := range numbers {
		numbers[i] = '0'
	}
	for i, d := range digits {
		numbers[i] = d
	}

	result := make([]string, 0)
	// O(4*4*4*4) = O(256)
	for _, l1 := range letters[numbers[0]] {
		for _, l2 := range letters[numbers[1]] {
			for _, l3 := range letters[numbers[2]] {
				for _, l4 := range letters[numbers[3]] {
					st := l1 + l2 + l3 + l4
					result = append(result, st)
				}

				if len(letters[numbers[3]]) == 0 {
					st := l1 + l2 + l3
					result = append(result, st)
				}
			}

			if len(letters[numbers[2]]) == 0 {
				st := l1 + l2
				result = append(result, st)
			}
		}
	}
	return result
}
