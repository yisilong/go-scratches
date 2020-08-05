package main

import (
	"fmt"
	"reflect"
)

func main() {
	var p *int
	v := reflect.ValueOf(p)
	fmt.Println(p, v)
}
