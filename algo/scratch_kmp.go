package main

import (
    "fmt"
)

// kmp算法
func GetNextArray(str string) []int {
    length := len(str)
    if length == 0 {
        return nil
    }
    nextArr := make([]int, length, length)
    nextArr[0] = -1
    cn := 0
    for pos := 2; pos < length; {
        if str[pos-1] == str[cn] {
            cn++
            nextArr[pos] = cn
            pos++
        } else if nextArr[cn] == -1 {
            nextArr[pos] = 0
            pos++
        } else {
            cn = nextArr[cn]
        }
    }
    return nextArr
}

func Search(src string, pattern string) int {
    nextArr := GetNextArray(pattern)
    var i, j int
    for i < len(src) && j < len(pattern) {
        if src[i] == pattern[j] {
            i++
            j++
        } else if nextArr[j] == -1 {
            i++
        } else {
            j = nextArr[j]
        }
    }
    if j == len(pattern) {
        return i - j
    }
    return -1
}

func main() {
    fmt.Println(Search("ABCDABCEAAAABASABCDABCADABCDABCEAABCDABCEAAABASABCDABCAABLAKABCDABABCDABCEAAADSFDABCADABCDABCEAAABCDABCEAAABASABCDABCADABCDABCEAAABLAKABLAKK", "ABCDABCEAAABASABCDABCADABCDABCEAAABLAK"))
}
