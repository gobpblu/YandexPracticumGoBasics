package main

import "fmt"

func main() {
    number := 10

    isMul := number%2 == 0 || number%5 == 0

    fmt.Println(isMul)
}