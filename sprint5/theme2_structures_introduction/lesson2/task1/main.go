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

type House struct {
	Name  string
	Rooms [][]Room
}

func main() {
	house := House{
		Name: "Дом v1",
		Rooms: [][]Room{
			{
				Room{
					Name: "Кладовка",
					Things: []Thing{
						Thing{Name: "топор", Weight: 3000},
						Thing{Name: "фонарик"},
						Thing{Name: "брелок"},
					},
				},
				Room{
					Name: "Котельная",
					Things: []Thing{
						Thing{Name: "верёвка", Weight: 200},
						Thing{Name: "рюкзак", Weight: 500},
					},
				},
			},
			{
				Room{
					Name: "Столовая",
					Things: []Thing{
						Thing{Name: "ручка"},
						Thing{Name: "кольцо"},
						Thing{Name: "карта"},
						Thing{Name: "бинт"},
					},
				},
			},
		},
	}
	fmt.Println(house)
}
