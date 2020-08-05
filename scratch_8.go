package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	m := md5.New()
	m.Write([]byte("hello"))
	var s string
	s = hex.EncodeToString(m.Sum(nil))
	fmt.Println(s)
	s = fmt.Sprintf("%x", m.Sum(nil))
	fmt.Println(s)
}
