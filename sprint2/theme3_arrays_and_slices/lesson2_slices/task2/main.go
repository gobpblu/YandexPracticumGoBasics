package main

import "fmt"

func main() {
	s := "Мезчхе%зцчхкьексцє%з%9%ьеце%ме%ширус%йусе"
	// расшифруйте сообщение и выведете его
	r := []rune(s)
	for i := range r {
		r[i] -= 5
	}
	s = string(r)
	fmt.Println(s)
}
