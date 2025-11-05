package main

import (
	"fmt"
	"unicode"
)

func upper(ch <-chan rune, done chan<- struct{}) {
	result := []rune{}
	for symbol := range ch {
		result = append(result, unicode.ToUpper(symbol))
	}
	fmt.Print(string(result))
	done <- struct{}{}
}

func main() {
	// определите буферизированный канал с рунами (rune)
	ch := make(chan rune, 1)
	done := make(chan struct{})

	go upper(ch, done)
	for _, v := range "Каждый охотник желает знать, где сидит фазан" {
		ch <- v
	}
	close(ch)
	<-done
}
