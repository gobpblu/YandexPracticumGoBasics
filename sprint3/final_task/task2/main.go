package main

import (
	"fmt"
	"sort"
)

// Median возвращает медиану числовой последовательности.
// Напишите код функции
// ...

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	medians := []int{
		0, 57, 35, 99, 4, 44, 43, 22000,
	}

	for i, list := range lists {
		if median := Median(list); median != medians[i] {
			fmt.Printf("median %d: %d != %d\n", i, medians[i], median)
		}
	}
	fmt.Println("Тестирование завершено")
}

func Median(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	if len(nums)%2 != 0 {
		return copyNums[len(copyNums)/2]
	} else {
		firstNum := copyNums[len(copyNums)/2]
		secondNum := copyNums[len(copyNums)/2-1]
		return (firstNum + secondNum) / 2
	}
}
