package main

import "fmt"

func main() {
	for i := 3; i < 10; i += 2 {
		for j := 3; j < 10; j += 2 {
			fmt.Printf("%dx%d = %d\n", i, j, i*j)
		}
	}
}
