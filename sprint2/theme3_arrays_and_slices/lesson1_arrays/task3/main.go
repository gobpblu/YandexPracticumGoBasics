package main

import "fmt"

func main() {
	// рекомендованные фильмы
	recommended := [...]string{"Хатико", "23", "Достучаться до небес",
		"Хакеры", "Трон", "1408"}
	// коллекция
	collection := [...]string{"Трон", "Военные игры", "Тихушники",
		"Джонни Мнемоник", "Хакеры", "Нирвана", "23", "Враг государства",
		"Взлом", "Пароль рыба-меч", "Сеть", "Кто я"}

	// допишите программу
	// ...
outsideLoop:
	for _, recommendEntry := range recommended {
		for _, entry := range collection {
			if entry == recommendEntry {
				fmt.Println(recommendEntry)
				continue outsideLoop
			}
		}
	}
}
