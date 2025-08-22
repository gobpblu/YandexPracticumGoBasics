package main

import "fmt"

func main() {
	bakery := []string{"Хлеб белый", "Хлеб Дарницкий", "Батон", "Булочка Домашняя", "Хачапури",
		"Сосиска в тесте", "Беляш", "Растегай", "Самса", "Пирожок с картошкой"}
	price := []int{43, 47, 52, 35, 63, 70, 82, 87, 91, 37}

	var result map[string]int // мапа с прайс-листом
	result = make(map[string]int, len(bakery))

	for i := range bakery {
		result[bakery[i]] = price[i]
	}

	fmt.Println(result["Беляш"], result["Пирожок с картошкой"], result["Батон"])
}
