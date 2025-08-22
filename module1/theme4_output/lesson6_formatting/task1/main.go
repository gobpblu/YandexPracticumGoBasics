package main

import "fmt"

func main() {
	oneHundred := 100
	rubles := "рублей"
	friends := "друзей"

	proverb := fmt.Sprintf("Не имей %d %s, а имей %d %s", oneHundred, rubles, oneHundred, friends)
	fmt.Println(proverb)
}