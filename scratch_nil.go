package main

import (
    "fmt"
)

type People interface {
    Show()
}

type Student struct{}

func (stu *Student) Show() {
}

func live1() People {
    return nil
}
func live2() People {
    var stu *Student
    return stu
}
func live3() *Student {
    return nil
}
func live4() *Student {
    var stu *Student
    return stu
}

// interface类型比较特殊，与nil比较会比较type和value是否同时为nil

func main() {
    fmt.Println(live1() == nil)  // true
    fmt.Println(live2() == nil)  // false
    fmt.Println(live3() == nil)  // true
    fmt.Println(live4() == nil)  // true
    var s *Student
    var p People
    fmt.Println(s == nil)   // true
    fmt.Println(p == nil)   // true
    p = nil
    fmt.Println(p == nil)   // true
    p = s
    fmt.Println(p == nil)   // false
}
