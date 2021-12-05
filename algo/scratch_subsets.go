package main

import (
	"fmt"
	"leetcode/helper"
)

// 子集
func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	result := [][]int{{}}
	for i := 0; i < length; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			tmp := len(result[j])
			one := make([]int, tmp+1, tmp+1)
			one[tmp] = nums[i]
			copy(one, result[j])
			result = append(result, one)
		}
	}
	return result
}

func testSubsets(in string) {
	nums := helper.ParseArray(in)
	fmt.Printf("%v  =>  %v\n", nums, subsets(nums))
}

func main() {
	testSubsets("[]")
	testSubsets("[1]")
	testSubsets("[1,3]")
	testSubsets("[1,3,5]")
	testSubsets("[1,3,5,9]")
}
