package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	if nums[0] > 0 || nums[length-1] < 0 {
		return nil
	}

	var results [][]int
	for i := 0; i < length-2; {
		if nums[i] > 0 {
			break
		}
		front := i + 1
		end := length - 1
		for front < end {
			sum := nums[i] + nums[front] + nums[end]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[front], nums[end]})
			}
			if sum <= 0 {
				for front++; front < end && nums[front] == nums[front-1]; front++ {
				}
			}
			if sum >= 0 {
				for end--; front < end && nums[end] == nums[end+1]; end-- {
				}
			}
		}
		for i++; i < length-2 && nums[i] == nums[i-1]; i++ {
		}
	}
	return results
}

func Print(nums []int) {
	fmt.Printf("%v: %v\n", nums, threeSum(nums))
}

func main() {
	Print(nil)
	Print([]int{1})
	Print([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	Print([]int{-2, 0, 0, 2, 2})
	Print([]int{-1, 0, 1, 2, -1, -4})
}
