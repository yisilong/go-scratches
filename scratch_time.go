package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    t22 := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.Local)
    fmt.Println("time", now, t22)
    if now.After(t22) {
        t22 = t22.AddDate(0, 0, 1)
    }
    fmt.Println("time", now, t22)
    leftTs := time.Duration(t22.Unix() - now.Unix())
    fmt.Println("left_ts", t22.Unix(), now.Unix(), time.Second*leftTs)
    start := t22.Add(-time.Hour*24)
    sid := fmt.Sprintf("%s-%s", start.Format("2006010215"), t22.Format("2006010215"))
    fmt.Println(sid)
}
