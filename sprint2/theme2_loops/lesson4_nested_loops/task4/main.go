package main

import "fmt"

func main() {
	volume := 5_000_000

top:
	for i := 100; i <= 300; i++ {
		for j := 100; j <= 300; j++ {
			if i*i*j == volume {
				fmt.Println(i, i, j, volume)
				break top
			}
		}
	}
}
