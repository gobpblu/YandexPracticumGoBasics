package main

import "fmt"

func main() {
	words := []string{"сервер", "cистема", "специалист", "слайc", "процессор",
		"масcив", "строка", "максимум", "cпоcоб", "парсер", "условие"}
	var mistakes []string

	for _, word := range words {
		wasCorrected := false
		newWord := []rune(word)
		for i, ch := range newWord {
			if ch == 'c' {
				newWord[i] = '?'
				wasCorrected = true
			}
		}
		if wasCorrected {
			mistakes = append(mistakes, string(newWord))
		}
	}

	fmt.Println(mistakes)
}
