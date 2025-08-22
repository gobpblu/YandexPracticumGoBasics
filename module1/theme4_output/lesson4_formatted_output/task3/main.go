package main

import "fmt"

func main() {
	fmt.Println("Алиса, выведи тип и значение для каждой переменной!\n")

	var number uint = 56
	floatNumber := 0.21442678
	str := "This is a string"
	isTrue := true

	// напишите ваш код ниже
	fmt.Printf("Тип: %T, значение: %d\n", number, number)
	fmt.Printf("Тип: %T, значение: %.3f\n", floatNumber, floatNumber)
	fmt.Printf("Тип: %T, значение: %q\n", str, str)
	fmt.Printf("Тип: %T, значение: %t\n", isTrue, isTrue)
}
