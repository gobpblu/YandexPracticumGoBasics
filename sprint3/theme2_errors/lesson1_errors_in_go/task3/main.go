package main

import (
	"errors"
	"fmt"
	"strconv"
)

// создаем переменную ошибки тут
var (
	ErrConvToFloat       = errors.New("ошибка преобразования строки в число с плавающей точкой")
	ErrNotPositiveNumber = errors.New("число должно быть больше нуля")
)

// ниже ваша функция сложения
func addPositive(num1, num2 string) (float64, error) {
	number1, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return 0, ErrConvToFloat
	}
	number2, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return 0, ErrConvToFloat
	}

	if number1 <= 0 || number2 <= 0 {
		return 0, ErrNotPositiveNumber
	}

	return number1 + number2, nil
}

func main() {
	num1 := "5.2"
	num2 := "2.7"

	result, err := addPositive(num1, num2)

	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения", result)
	}

	num1 = "two"
	num2 = "five"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}

	num1 = "5.4"
	num2 = "0.0"

	result, err = addPositive(num1, num2)
	if err != nil {
		err = fmt.Errorf("ошибка в ходе выполнения программы: %v", err)
		fmt.Println(err)
	} else {
		fmt.Println("результат сложения:", result)
	}
}
