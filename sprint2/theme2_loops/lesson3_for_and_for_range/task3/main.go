package main

import (
	"fmt"
	"strings"
)

func main() {
	var hidden string // результирующая строка со звёздочками
	email := `vasyapupkin33@mail.ru`

	index := strings.Index(email, "@")
	if index == -1 || index < 3 {
		hidden = email
	} else {
		count := 0
		for i, ch := range email {
			if count < 2 || i >= index {
				hidden += string(ch)
			} else {
				hidden += "*"
			}
			count++
		}
	}
	fmt.Println(hidden)
}
