package main

import (
    "fmt"
    "sync"
)

var list []int
var wg sync.WaitGroup
var mu sync.Mutex

func do() {
	mu.Lock()
    for i := 0; i < 100; i++ {
        list = append(list, i)
    }
	mu.Unlock()
	wg.Done()
}

func main() {
	wg.Add(8)
    for i := 0; i < 8; i++ {
        go do()
    }
	wg.Wait()
    // проверка содержимого у слайса
    sum := 0
    for _, v := range list {
        sum += v
    }
    fmt.Println(len(list), sum)
}