package main

import "fmt"

func DivSizes(x, y int) (int, int) {
    return x / 2, y / 2 
}

func main() {
    x, y := DivSizes(10, 20)
    fmt.Println(x, y)
}