package main

import "fmt"

func main() {
    // исходный массив
    numbers := [...]int{1, 3, 8, 19, 42}
    
    // 1. Создайте и заполните слайс указателей на элементы массива
    // 2. Пройдитесь по слайсу и умножьте на три все значения, на которые
    //    ссылаются указатели
    // ...
	pNumbers := []*int{&numbers[0], &numbers[1], &numbers[2], &numbers[3], &numbers[4] }
	for _, p := range pNumbers {
		*p *= 3
	}

    fmt.Println(numbers)
}