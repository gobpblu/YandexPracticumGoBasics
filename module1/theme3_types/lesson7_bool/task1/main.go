package main

import "fmt"

func main() {
	minAge := 18
	maxAge := 30

	age := 35

	isValidAge := age >= minAge && age <= maxAge

	fmt.Println(isValidAge)
}
