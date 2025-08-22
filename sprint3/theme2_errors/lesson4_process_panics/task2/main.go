package main

import (
	"fmt"
)

func panicGeneration() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановление в panicGeneration()")
		}
	}()
	panic("Паника в panicGeneration()")
}

func main() {
	panicGeneration()

	fmt.Println("Строка после паники")
}