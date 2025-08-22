package main

import (
	"fmt"
	"sort"
)

func main() {
	// прайс-лист выпечки
	bakery := map[string]int{
		"Хлеб белый":          43,
		"Хлеб Дарницкий":      47,
		"Батон":               52,
		"Булочка Домашняя":    35,
		"Хачапури":            63,
		"Сосиска в тесте":     70,
		"Беляш":               82,
		"Растегай":            87,
		"Самса":               91,
		"Пирожок с картошкой": 37,
	}

	sorted := make([]string, 0, len(bakery))
	for name := range bakery {
		sorted = append(sorted, name)
	}

	sort.Strings(sorted)

	for _, sortName := range sorted {
		fmt.Println(sortName, ":", bakery[sortName])
	}
}
