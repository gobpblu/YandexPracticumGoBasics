package main

import "fmt"

func main() {
	i := 0

	for i != 20 {
		switch i {
		case 5:
			break
		case 10:
		case 15:
			i += 5
		}
		i++
	}
	fmt.Print(i)
}
