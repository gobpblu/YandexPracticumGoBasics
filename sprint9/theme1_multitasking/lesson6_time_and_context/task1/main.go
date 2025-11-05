package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(200 * time.Millisecond)
	timer := time.NewTimer(4 * time.Second)

	var i = 1
	for {
		select {
		case <-ticker.C:
			if i%5 == 0 {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
			i++
		case <-timer.C:
			return
		}
	}
}
