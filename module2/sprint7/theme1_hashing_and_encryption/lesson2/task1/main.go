package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Winner(rnd *rand.Rand, bets map[int]string) string {
    var minDiff float64 = 100.0
	generatedNum := rnd.Int63() % 100
	winner := ""

	for bet, name := range bets {
		diff := math.Abs(float64(generatedNum - int64(bet)))
		if diff < minDiff {
			minDiff = diff
			winner = name
		}
	}

	return winner
}

func main() {
    rnd := rand.New(rand.NewSource(1001))
    // кто какое число загадал
    bets := map[int]string{
        20: "Маша",
        34: "Игорь",
        77: "Олег",
        51: "Света",
        2:  "Саша",
    }

    fmt.Println("Победитель:", Winner(rnd, bets))
}