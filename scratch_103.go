package main

import "fmt"

func maxArea(height []int) int {
	length := len(height)
	if length < 1 {
		return 0
	}

	var max, curr int
	var l int
	r := length - 1
	for l < r {
		if height[l] < height[r] {
			curr = height[l] * (r - l)
			l++
		} else {
			curr = height[r] * (r - l)
			r--
		}
		if curr > max {
			max = curr
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
