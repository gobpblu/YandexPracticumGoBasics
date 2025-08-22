package main

import (
	"fmt"
	"errors"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("нельзя делить на ноль")
	}
    
    return x/y, nil
}

func main() {
    var a, b int
    fmt.Println(div(a, b))
    a,b = 27,3
    fmt.Println(div(a, b))
}