package main

import "fmt"

func printLastBand(bands []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Printf("Последняя добавленная музыкальная группа %s мне нравится больше всего", bands[len(bands)])
}

func main() {
	bands := []string{
		"30 Seconds To Mars",
		"Garbage",
		"The Cardigans",
	}

	printLastBand(bands)

	fmt.Println("Паники нет, таблетка действует!")
}
