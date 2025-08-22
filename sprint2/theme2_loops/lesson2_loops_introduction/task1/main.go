package main

import "fmt"

func main() {

	count := 0 // количество отжиманий

	i := 1 // счётчик дней

	// добавьте цикл for с подсчётом отжиманий
	// ...
	for {
		count += i
		if i == 100 {
			break
		}
		i++
	}
	fmt.Println("Лёха сделает", count, "отжиманий за 100 дней")
}
