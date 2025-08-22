package main

import "fmt"

func main() {
	count := 2 // количество бактерий
	day := 1   // счётчик дней

	// реализуйте for с условием и подсчётом бактерий в теле цикла
	// ...
	for day <= 30 {
		if day > 10 {
			count -= count / 10
		}
		count += count / 2
		day++
	}

	fmt.Println("Количество бактерий через 30 дней:", count)
}
