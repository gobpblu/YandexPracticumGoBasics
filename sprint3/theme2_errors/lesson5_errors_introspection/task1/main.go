package main

import (
	"errors"
	"fmt"
)

// создаем новую ошибку
var ErrWrongArgument = errors.New("incorrect second argument")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("[div(%d, %d)] %w: can't divide %d by zero", a, b, ErrWrongArgument, a)
	}
	return a / b, nil
}

func main() {
	res, err := div(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
