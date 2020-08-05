package main

import (
	"fmt"
	"strconv"
)

func FloatToInt(v float64) int64 {
	s := fmt.Sprintf("%.0f", v)
	value, _ := strconv.ParseInt(s, 10, 0)
	return value
}

func main() {
	fmt.Println(int(float32(1.0)))
	fmt.Println(int(FloatToInt(1.5)))
}
