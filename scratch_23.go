package main

import (
	"fmt"
)
var T int64 = a()
var p map[string]int = c()

func init() {
	fmt.Println("init in main.go")
}

func a() int64 {
	fmt.Println("calling a()")
	return 2
}

func c() map[string]int {
	return nil
}

func main() {
	fmt.Println("calling main")

}


