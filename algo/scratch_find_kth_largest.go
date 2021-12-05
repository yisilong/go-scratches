package main

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := partation(nums, left, right)
		if mid == k-1 {
			return nums[mid]
		} else if mid > k-1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1 // can not arrive
}

func partation(nums []int, left, right int) int {
	mid := (left + right) / 2
	nums[left], nums[mid] = nums[mid], nums[left]
	l, r := left+1, right
	pivot := nums[left]
	for l <= r {
		for l < r && nums[l] >= pivot {
			l++
		}
		for nums[r] < pivot {
			right--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		} else {
			break
		}
	}
	nums[left], nums[r] = nums[r], nums[left]
	return r
}

func main() {

}
