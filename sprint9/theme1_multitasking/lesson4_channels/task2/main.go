package main

import (
	"fmt"
)

func do(in, out chan int) {
	defer close(out)

	for v := range in {
		out <- v * 2
	}
}

func main() {
	chIn := make(chan int)
	chOut := make(chan int)

	// запустите горутину, которая преобразует числа
	go do(chIn, chOut)

	// запустите горутину, которая отправляет числа на обработку
	go func() {
		defer close(chIn)

		for i := range 51 {
			chIn <- i
		}
	}()

	// в цикле читайте числа из результирующего канала. Используйте for range
	for v := range chOut {
		fmt.Print(v, " ")
	}
}
