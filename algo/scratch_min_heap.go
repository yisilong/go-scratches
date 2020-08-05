package main

import (
    "fmt"
)

// 小顶堆
type MinHeap struct {
    size int
    arr  []int
}

// 下沉
func (h *MinHeap) Sink(i int) {
    for {
        left := i*2 + 1
        right := left + 1
        child := i
        if left < h.size && h.arr[left] < h.arr[i] {
            child = left
        }

        if right < h.size && h.arr[right] < h.arr[child] {
            child = right
        }

        if child == i {
            break
        }
        h.arr[child], h.arr[i], i = h.arr[i], h.arr[child], child
    }
}

// 上浮
func (h *MinHeap) Swim(i int) {
    for i != 0 {
        parent := (i - 1) / 2
        if h.arr[i] >= h.arr[parent] {
            break
        }
        h.arr[i], h.arr[parent], i = h.arr[parent], h.arr[i], parent
    }
}

// 堆化
func (h *MinHeap) Heapify() {
    for i := h.size/2 - 1; i >= 0; i-- {
        h.Sink(i)
    }
}

func (h *MinHeap) Top() int {
    var v int
    if h.size > 0 {
        v = h.arr[0]
    }
    return v
}

func (h *MinHeap) Pop() int {
    var v int
    if h.size > 0 {
        v = h.arr[0]
        if h.size > 1 {
            h.arr[0] = h.arr[h.size-1]
            if h.size > 2 {
                h.Sink(0)
            }
        }
        h.size--
    }
    return v
}

func (h *MinHeap) Sort() []int {
    var arr []int
    if h.size > 0 {
        arr = make([]int, h.size, h.size)
        for h.size > 0 {
            arr[h.size-1] = h.arr[0]
            if h.size > 1 {
                h.arr[0] = h.arr[h.size-1]
                if h.size > 2 {
                    h.Sink(0)
                }
            }
            h.size--
        }
    }
    return arr
}

func (h *MinHeap) Add(v int) {
    if h.size == len(h.arr) {
        if v > h.arr[0] {
            h.arr[0] = v
            h.Sink(0)
        }
    } else {
        h.arr[h.size] = v
        h.size++
        h.Swim(h.size - 1)
    }
}

func (h *MinHeap) Insert(v int) bool {
    if h.size == len(h.arr) {
        return false
    }

    h.arr[h.size] = v
    h.size++
    return true
}

func NewMinHeap(cap int) *MinHeap {
    h := &MinHeap{}
    h.arr = make([]int, cap, cap)
    return h
}

func main() {
    h := NewMinHeap(14)
    for _, v := range []int{99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92} {
        h.Insert(v)
    }
    h.Heapify()
    fmt.Println(h.Sort())
}
