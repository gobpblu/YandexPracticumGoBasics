package main

import "fmt"

func main() {
	fmt.Printf("Dec\tChar\tHex\n")
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\t%c\t%x\n", i, i, i)
	}
}
