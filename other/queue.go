package other

// leetcode: 1375
func numTimesAllBlue(light []int) int {
	// original
	// queue := make([]int, 0)
	// leftEnabled := 0
	// counter := 0
	// for _, el := range light {
	// 	if leftEnabled != el-1 {
	// 		queue = append(queue, el)
	// 		continue
	// 	}
	// 	leftEnabled = el
	// 	lastInd := 0
	// 	sort.Ints(queue)
	// 	for i := 0; i < len(queue); i++ {
	// 		if queue[i]-1 != leftEnabled {
	// 			lastInd = i
	// 			break
	// 		}
	// 		leftEnabled = queue[i]
	// 		lastInd = i + 1
	// 	}
	// 	queue = queue[lastInd:]
	// 	if len(queue) == 0 {
	// 		counter++
	// 	}
	// }
	// return counter

	// optimal
	counter := 0
	maxVal := 0
	for i, v := range light {
		if v > maxVal {
			maxVal = v
		}
		// 5 bulb cant be enabled before 5 step
		if maxVal == i+1 {
			counter++
		}
	}
	return counter
}
