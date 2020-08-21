package main

import "fmt"

// 整数反转
func reverse(x int) int {
	const SENTRY = (1 << 31) / 10
	var result int
	for x != 0 {
		num := x % 10
		x /= 10
		if result > SENTRY || result == SENTRY && num > 7 {
			return 0
		}
		if result < -SENTRY || result == -SENTRY && num < -8 {
			return 0
		}
		result = result*10 + num
	}
	return result
}

func main() {
	fmt.Println(reverse(10))
	fmt.Println(reverse(-10))
	fmt.Println(reverse(12))
	fmt.Println(reverse(-12))
	fmt.Println(reverse(1005))
	fmt.Println(reverse(-1005))
}
