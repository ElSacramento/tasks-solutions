package other

// leetcode: 929
// O(n*m) + space O(n)
func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, em := range emails {
		key := getRealEmail(em)
		set[key] = struct{}{}
	}
	return len(set)
}

func getRealEmail(s string) string {
	sepInd := -1
	ignoreAll := false
	result := make([]rune, 0, len(s))
	for i, ch := range s {
		if ch == '@' {
			sepInd = i
			result = append(result, ch)
			ignoreAll = false
			continue
		}
		// login part
		if sepInd == -1 {
			if ignoreAll {
				continue
			}
			if ch == '.' {
				// ignore
				continue
			}
			if ch == '+' {
				ignoreAll = true
				continue
			}
		}
		result = append(result, ch)
	}
	return string(result)
}
