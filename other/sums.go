package other

import (
	"strconv"
)

// O(n)
// leetcode: 1010
func numPairsDivisibleBy60(time []int) int {
	counter := 0
	// remainder: number of elements with such remainder
	remainders := make(map[int]int, 60)

	for _, t := range time {
		if t%60 == 0 { // a%60 = 0 , b%60 = 0
			counter += remainders[0]
		} else { // a%60 + b%60 = 60
			counter += remainders[60-t%60]
		}
		remainders[t%60] += 1
	}

	return counter
}

func count(nums []int, start, end, x, counter int) int {
	if x-nums[start] == 0 {
		return counter + 1
	}
	if x-nums[end] == 0 {
		return counter + 1
	}

	if x-nums[start] > 0 {
		counter += count(nums, start+1, end, x-nums[0], counter)
	}
	if x-nums[end] > 0 {
		counter += count(nums, start, end-1, x-nums[end], counter)
	}
	return counter
}

// todo: find task description
func minOperations(nums []int, x int) int {
	counter := 0
	lastInd := len(nums) - 1

	if x-nums[0] == 0 {
		return 1
	}
	if x-nums[lastInd] == 0 {
		return 1
	}

	if x-nums[0] < 0 && x-nums[lastInd] < 0 {
		return -1
	}

	if x-nums[0] > 0 {
		counter += count(nums, 1, lastInd, x-nums[0], counter)
	}
	if x-nums[lastInd] > 0 {
		counter += count(nums, 0, lastInd-1, x-nums[lastInd], counter)
	}

	return counter
}

// leetcode: 682
// simple calculation O(n) + space for sum on each step
func calPoints(ops []string) int {
	points := make([]int, len(ops))

	prevInd := -1
	for _, s := range ops {
		switch s {
		case "+":
			points[prevInd+1] = points[prevInd] + points[prevInd-1]
			prevInd++
		case "D":
			points[prevInd+1] = 2 * points[prevInd]
			prevInd++
		case "C":
			points[prevInd] = 0
			prevInd--
		default:
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			points[prevInd+1] = val
			prevInd++
		}
	}

	sum := 0
	for _, v := range points {
		sum += v
	}
	return sum
}
