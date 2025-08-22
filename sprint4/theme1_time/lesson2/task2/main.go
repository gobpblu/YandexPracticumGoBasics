package main

import (
	"fmt"
	"time"
)

func main() {
	// создайте переменные типа time.Time
	v1_20 := time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)
	v1_21 := time.Date(2023, time.August, 8, 0, 0, 0, 0, time.UTC)
	// получите интервал между датами
	duration := v1_21.Sub(v1_20)
	// вычислите количество дней
	days := int(duration.Hours() / 24)
	fmt.Printf("Между выходами версий прошло %d дней", days)
}
