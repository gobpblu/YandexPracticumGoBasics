package main

import "fmt"

func main() {
	// ассортимент выпечки
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
	// список покупок
	buy := []string{"Растегай", "Самса", "Хачапури", "Хлеб Подовый", "Беляш"}
	// вывод общего количества позиций в bakery
	fmt.Println("Всего позиций:", len(bakery))

	countInCart, mistakesCount, sumToPay := 0, 0, 0
	for _, buyPrice := range buy {
		price, ok := bakery[buyPrice]
		if !ok {
			mistakesCount++
		} else {
			sumToPay += price
			countInCart++
		}
	}

	// вставьте вторым параметром соответствующие переменные
	fmt.Println("Ошибок:", mistakesCount)  // количество ошибочных товаров
	fmt.Println("В корзине:", countInCart) // количество корректных товаров
	fmt.Println("К оплате:", sumToPay)
}
