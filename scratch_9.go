package main

import (
	"encoding/json"
	"fmt"
)

type St struct {
	Name string
}

func main() {
	s := St{Name:"ysl"}
	d , err := json.Marshal(&s)
	fmt.Printf("%v, %v\n", d, err)
	fmt.Println(d, err)
}
