package main

import (
	"fmt"
	"sort"
)

// Mode возвращает моды числовой последовательности.
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
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{-7, 102}, {},
		{20000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}
	for i, list := range lists {
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
		}
	}
	fmt.Println("Тестирование завершено")
}

func Mode(originalNums []int) ([]int, int) {
	nums := make([]int, 0)
	if len(originalNums) == 0 {
		return nums, 1
	}
	// copy(nums, originalNums)
	// sort.Ints(nums)
	var maxValue int
	uniqueNums := make(map[int]int, 0)
	for _, num := range originalNums {
		_, ok := uniqueNums[num]
		if !ok {
			uniqueNums[num] = 1
		} else {
			uniqueNums[num]++
		}
		if uniqueNums[num] > maxValue {
			maxValue = uniqueNums[num]
		}
	}

	for key, value := range uniqueNums {
		if value == maxValue && maxValue > 1 {
			nums = append(nums, key)
		}
	}
	sort.Ints(nums)
	return nums, maxValue
}
