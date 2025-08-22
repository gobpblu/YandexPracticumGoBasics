package main

import "fmt"

type Thing struct {
    Name   string
    Weight int
}

type Room struct {
    Name   string
    Things []Thing
}

func main() {
    var pen Thing

    pen.Name = "ручка"
    pen.Weight = 50

    // 1. Создайте переменную room типа Room
    // 2. Присвойте ей имя "Библиотека"
    // 3. Добавьте в слайс Things ручку pen
    room := Room {}
	room.Name = "Библиотека"
	room.Things = append(room.Things, pen)

    fmt.Println(room)
}