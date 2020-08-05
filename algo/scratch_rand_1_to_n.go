package main

import (
    "fmt"
    "math/rand"
)

// [1, m] 生成 [1, n]

func rand1ToM(m int) int {
    return rand.Intn(m) + 1
}

func rand1ToN(m, n int) int {
    mSysNum := toMSysNum(n-1, m)
    randNum := randMSysNumLessN(mSysNum, m)
    return fromMSysNum(randNum, m) + 1
}

type randFunc func() int

func randMM(f randFunc, m int) randFunc {
    return func() int {
        fmt.Println("m----", m)
        return f()*m + f()
    }
}

func rand1ToNv2(m, n int) int {
    m1 := m
    n1 := n
    f := func() int { return rand1ToM(m) - 1 }
    for n1 > m1 {
        f = randMM(f, m1)
        m1 *= m1
    }

    n1 = m1 - m1%n

    num := n1
    for num >= n1 {
        num = f()
    }

    return num%n + 1
}

func toMSysNum(v, m int) []int {
    var mSysNum []int
    for v != 0 {
        mSysNum = append(mSysNum, v%m)
        v = v / m
    }
    return mSysNum
}

func randMSysNumLessN(nMSys []int, m int) []int {
    length := len(nMSys)
    if length == 0 {
        return nil
    }

    randMSys := make([]int, length, length)
    i := length
    needCheck := true
    for i > 0 {
        i--
        randMSys[i] = rand1ToM(m) - 1
        if needCheck {
            if randMSys[i] > nMSys[i] {
                i = length
                continue
            } else if randMSys[i] < nMSys[i] {
                needCheck = false
            }
        }
    }
    return randMSys
}

func fromMSysNum(mSysNum []int, m int) int {
    var num int
    for i := len(mSysNum); i > 0; i-- {
        num = num*m + mSysNum[i-1]
    }
    return num
}

func main() {
    m := make(map[int]int)
    for i := 0; i < 100000; i++ {
        n := rand1ToN(5, 5)
        m[n] = m[n] + 1
    }

    m2 := make(map[int]int)
    for i := 0; i < 100000; i++ {
        n := rand1ToNv2(5, 5)
        m2[n] = m2[n] + 1
    }
    fmt.Println(m)
    fmt.Println(m2)
}
