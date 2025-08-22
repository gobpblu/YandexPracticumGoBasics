package main

import "fmt"

func main() {
	track1 := 189
	track2 := 148
	track3 := 210
	track4 := 143

	// напишите свой код ниже
	totalDuration := track1 + track2 + track3 + track4
	minutes := totalDuration / 60
	seconds := totalDuration % 60

	musicDurationMessage := fmt.Sprintf("Вы слушали музыку %d минут %d секунд", minutes, seconds)
	fmt.Println(musicDurationMessage)
}
