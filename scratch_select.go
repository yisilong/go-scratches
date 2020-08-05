package main

import (
    "fmt"
    "time"
)

/*
1. 除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
2. 除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
3. 如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
4. 如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞
*/
func doSelect(ch <-chan int, c chan int) {
    fmt.Println("start select")
    for i := 0; i < 20; i++ {
        select {
        case <-ch:
            fmt.Println("ch -----------")
        default:
            fmt.Println("default -----------")
        }
        time.Sleep(30 * time.Millisecond)
    }
    fmt.Println("end select")
    close(c)
}

func product(ch chan<- int) {
    for {
        <-time.After(30 * time.Millisecond)
        ch <- 3
    }
}

func main() {
    ch, done := make(chan int), make(chan int)
    go doSelect(ch, done)
    go product(ch)
    <-done
}
