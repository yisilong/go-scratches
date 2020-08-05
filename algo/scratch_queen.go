package main

import (
    "fmt"
)

// N 皇后
func queen(n int) int {
    if n < 1 {
        return 0
    }
    shuArr := make([]bool, n, n)
    pieArr := make([]bool, 2*n-1, 2*n-1)
    naArr := make([]bool, 2*n-1, 2*n-1)
    record := make([]int, n, n)
    return process1(0, n, record, shuArr, pieArr, naArr)
}

func process1(row, n int, record []int, shuArr, pieArr, naArr []bool) int {
    if row == n {
        fmt.Println("got one:", record)
        return 1
    }

    num := 0
    for col := 0; col < n; col++ {
        pie := row + col
        na := n - 1 - row + col
        if shuArr[col] || pieArr[pie] || naArr[na] {
            continue
        }
        record[row] = col
        shuArr[col] = true
        pieArr[pie] = true
        naArr[na] = true
        num += process1(row+1, n, record, shuArr, pieArr, naArr)
        shuArr[col] = false
        pieArr[pie] = false
        naArr[na] = false
    }
    return num
}

func isValid(record []int, row, col int) bool {
    for i := 0; i < row; i++ {
        if col == record[i] || record[i]-col == row-i || record[i]-col == i-row {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(queen(13))
}
