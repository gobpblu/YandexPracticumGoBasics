package main

import "fmt"

type Hero struct {
	Name   string // имя героя
	Level  int    // уровень героя
	Health int    // здоровье героя
}

// IncLevel увеличивает уровень героя на единицу.
// добавьте метод IncLevel()
func (hero *Hero) IncLevel() {
	hero.Level++
}

// ChangeHealth изменяет здоровье героя на указанную величину.
// добавьте метод ChangeHealth(dif int)
func (hero *Hero) ChangeHealth(dif int) {
	hero.Health += dif
}

func main() {
	hero := Hero{
		Name:   "Буратино",
		Level:  10,
		Health: 35,
	}
	hero.IncLevel()
	hero.ChangeHealth(10)

	fmt.Println(hero)
}
