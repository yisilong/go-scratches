package main

import "fmt"

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

func Print(nums []int) {
	fmt.Printf("%v  =>  %v\n", nums, subsets(nums))
}

func main() {
	Print(nil)
	Print([]int{1})
	Print([]int{1, 3})
	Print([]int{1, 3, 5})
	Print([]int{1, 3, 5, 9})
}
