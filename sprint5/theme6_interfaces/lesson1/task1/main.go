package main

import "fmt"

// Character содержит общую информацию о герое.
type Character struct {
    Name   string
    Health int
}

// Warrior описывает воина.
type Warrior struct {
    Character
    Power int
}

// Mage описывает мага.
type Mage struct {
    Character
    Magic int
}

func (w Warrior) String() string {
	return fmt.Sprint("* ", w.Name)
}

func (m Mage) String() string {
	return fmt.Sprint("# ", m.Name)
}

func main() {
    // проверяем работу метода String()
    fmt.Println(Warrior{Character: Character{Name: "Крестоносец"}})
    fmt.Println(Warrior{Character: Character{Name: "Коммандос"}})
    fmt.Println(Mage{Character: Character{Name: "Шаман"}})
    fmt.Println(Mage{Character: Character{Name: "Друид"}})
}