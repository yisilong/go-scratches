package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fmt.Sprintf("%f", 10.900100))
	fmt.Println(fmt.Sprintf("%.2f", 10.900100))
	fmt.Println(strconv.FormatFloat(10.900100, 'f', -1, 64))
}
