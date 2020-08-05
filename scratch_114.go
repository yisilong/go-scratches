package main

import (
    "fmt"
)

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func min(num1, num2 int) int {
    if num1 < num2 {
        return num1
    } else {
        return num2
    }
}

func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}

func IfElse(b bool, one, two int) int {
    if b {
        return one
    } else {
        return two
    }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1, l2 := len(nums1), len(nums2)
    if l1 > l2 {
        return findMedianSortedArrays(nums2, nums1)
    }

    var lMax1, lMax2, rMin1, rMin2, c1, c2 int
    left, right := 0, 2*l1
    for left <= right {
        c1 = (left + right) / 2
        c2 = l1 + l2 - c1
        lMax1 = IfElse(c1 == 0, IntMin, nums1[(c1-1)/2])
        rMin1 = IfElse(c1 == 2*l1, IntMax, nums1[c1/2])
        lMax2 = IfElse(c2 == 0, IntMin, nums2[(c2-1)/2])
        rMin2 = IfElse(c2 == 2*l2, IntMax, nums2[c2/2])

        if lMax1 > rMin2 {
            right = c1 - 1
        } else if lMax2 > rMin1 {
            left = c1 + 1
        } else {
            break
        }
    }
    num := float64(max(lMax1, lMax2))
    num += float64(min(rMin1, rMin2))
    return num / 2.0
}

func Print(num1 []int, num2 []int) {
    fmt.Println(findMedianSortedArrays(num1, num2))
}

func main() {
    Print([]int{1, 2}, []int{3, 4})
}
