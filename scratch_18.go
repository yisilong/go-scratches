package main

import (
    "fmt"
)

// kmp算法
func GetNextArray(str string) []int {
    var nextArr []int
    length := len(str)
    if length > 0 {
        nextArr = make([]int, length, length)
        nextArr[0] = -1
        cnt := 0
        for i := 2; i < length; {
            if str[i-1] == str[cnt] {
                nextArr[i] = cnt + 1
                cnt++
                i++
            } else if nextArr[cnt] == -1 {
                // cn = 0
                // nextArr[i] = 0
                i++
            } else {
                cnt = nextArr[cnt]
            }
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
            j = 0 // compat nextVal
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
