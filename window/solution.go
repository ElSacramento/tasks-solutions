package window

// sliding window O(n)
func TotalFruit(tree []int) int {
	var (
		lastSwitch        = -1
		first, second     = tree[0], -1
		counter, maxCount = 1, 1
	)

	for i := 1; i < len(tree); i++ {
		prevValue := tree[i-1]
		value := tree[i]

		if prevValue == value {
			counter++
			continue
		}

		// value switch
		if second == -1 {
			second = value
			lastSwitch = i
			counter++
			continue
		}

		// second is already set
		if value == first || value == second {
			lastSwitch = i
			counter++
			continue
		}

		// new value
		if counter > maxCount {
			maxCount = counter
		}
		counter = i - lastSwitch + 1
		first = prevValue
		second = value
		lastSwitch = i
	}

	if counter > maxCount {
		maxCount = counter
	}

	return maxCount
}
