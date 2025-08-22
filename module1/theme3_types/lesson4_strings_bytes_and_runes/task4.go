package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Алиса, в какой папке находятся мои фото?"

	path := "C:\\Documents\\Photos"

	lowLetterPath := strings.ToLower(path)

	c := "c"
	d := "d"

	count := 1

	newPath := strings.Replace(lowLetterPath, c, d, count)

	fmt.Println(message)

	fmt.Println("path =", path)

	fmt.Println("lowLetterPath =", lowLetterPath)

	fmt.Println("newPath =", newPath)
}
