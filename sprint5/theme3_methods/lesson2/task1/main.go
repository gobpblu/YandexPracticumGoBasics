package main

import (
	"fmt"
)

// Hero описывает главного героя игры
type Hero struct {
	Name   string
	Health int
	Damage int
	Def    int
}

func (hero Hero) Defense() {
	fmt.Println(hero.Name, "заблокировал", hero.Def, "единиц урона")
}

// ниже напишите метод Special() для структуры Hero
func (hero *Hero) Special(healthPoints int) {
	fmt.Println("Количество здоровья было", hero.Health)
	hero.Health += healthPoints
	fmt.Println("Количество здоровья стало", hero.Health)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100, Damage: 30, Def: 20}
	myHero.Defense()
	myHero.Special(30)
}
