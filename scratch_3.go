package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	results := make([]int, length)
	k := 1
	for i := 0; i < length; i++ {
		results[i] = k
		k *= nums[i]
	}

	k = 1
	for i := length - 1; i >= 0; i-- {
		results[i] *= k
		k *= nums[i]
	}
	return results
}

func main() {

}
