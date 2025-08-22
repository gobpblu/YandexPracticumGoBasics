package main

import "fmt"

func main() {
	list := []int{1, 5, 4, 8, 7, 6, 0, 2, 9}
	var (
		index  int
		isLoop bool // должна быть true, если есть петля
		count  int
	)
	// добавьте в код поиск петли
	for index >= 0 && index < len(list) && count <= len(list) {
		index = list[index]
		count++
	}
	isLoop = count > len(list)
	fmt.Println("В последовательности имеется петля?", isLoop)
}
