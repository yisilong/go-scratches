package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello !!!!世界!!!!!!!!"
	s = strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", s)
	s = strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", s)
	s = strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", s)
}
