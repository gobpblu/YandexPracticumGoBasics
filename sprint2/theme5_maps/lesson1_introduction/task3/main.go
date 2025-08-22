package main

import "fmt"

func main() {
	// оценки по отдельным ученикам
	marks := map[string][]int{
		"Светлана":  {5, 4, 5, 5, 4},
		"Артём":     {3, 4, 4, 5, 3},
		"Александр": {2, 3, 3, 4},
		"Ольга":     {5, 5, 4, 4},
		"Мария":     {4, 3, 4, 4, 3, 5},
	}
	student := "Светлана"
	var average float32 // средний балл

	marks[student] = append(marks[student], 5)

	for _, mark := range marks[student] {
		average += float32(mark)
	}

	average = average / float32(len(marks[student]))

	fmt.Printf("%.2f\n", average)
}
