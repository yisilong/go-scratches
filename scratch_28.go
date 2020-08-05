package main

import (
    "fmt"
    "time"
)

func consumer(name string, ch <-chan int) {
    for {
        fmt.Println("consumer before-", name)
        select {
        case i := <-ch:
            fmt.Println("consumer after-", name, i)
        }
    }
}

func producer(name string, ch chan<- int, start int) {
    for i := start; i < start+5; i++ {
        fmt.Println("producer before-", name, i)
        ch <- i
        fmt.Println("producer after-", name, i)
    }
}

func main() {
    ch := make(chan int, 0)
    go consumer("1", ch)
    go consumer("2", ch)
    go consumer("3", ch)
    go producer("1", ch, 100)
    go producer("2", ch, 105)
    time.Sleep(5 * time.Second)
}
