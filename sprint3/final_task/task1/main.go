package main

import (
	"fmt"
	"math"
)

// Average возвращает среднее арифметическое элементов слайса []int.
// Напишите код функции
// ...

// Range возвращает размах числовой последовательности.
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
	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	ranges := []int{
		0, 0, 85, 200, 1, 109, 123, 80000,
	}

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
		}
		if r := Range(list); r != ranges[i] {
			fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
		}
	}
	fmt.Println("Тестирование завершено")
}

func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	var avg int = 0
	for _, num := range nums {
		avg += num
	}
	return float64(avg) / float64(len(nums))
}

func Range(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	return max - min
}
