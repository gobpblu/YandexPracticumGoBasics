package main

import "fmt"

type Character struct {
    Name   string // имя
    Health int    // здоровье
    Speed  int    // скорость
    Power  int    // сила
    Woman  bool   // true, если женский персонаж
}

// добавьте метод Say
// ...

func main() {
    ken := Character{
        Name: "Ken",
        Health: 100,
        Speed: 500,
        Power: 1000,
    }
    ken.Say("Привет! Как тебя зовут?")
}

func (characater Character) Say(msg string) {
	fmt.Printf("%s: %s", characater.Name, msg)
}