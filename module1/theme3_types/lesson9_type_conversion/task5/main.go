package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "21"
	y := "25.4"
	girlfriendAge, err := strconv.Atoi(x)
	myAge, err := strconv.ParseFloat(y, 64)

	if err != nil {
		panic(err)
	}

	sumOfAges := girlfriendAge + int(myAge)

	fmt.Println(sumOfAges)
}
