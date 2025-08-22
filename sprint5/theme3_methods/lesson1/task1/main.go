package main

import (
	"fmt"
)

type Hero struct {
	Name     string
	Health   int
	Location string
}

func (hero *Hero) MoveTo(location string) {
	hero.Location = location
	fmt.Println(hero.Name, "переместился в", location)
}

func main() {
	myHero := Hero{Name: "Артур", Health: 100}
	myHero.MoveTo("тронный зал")
}
