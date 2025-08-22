package main

import "fmt"

func main() {
	// оценки по отдельным ученикам
	marks := [][]int{
		{5, 4, 5, 5},
		{3, 4, 4, 5, 3},
		{2, 3, 3},
		{5, 5, 4},
		{4, 3, 4, 4, 3},
	}
	var average float32 // итоговый средний балл
	var count int

	for i := range marks {
		count += len(marks[i])
		for j := range marks[i] {
			average += float32(marks[i][j])
		}
	}

	average = average / float32(count)

	fmt.Printf("%.2f", average)
}
