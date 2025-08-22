package main

import "fmt"

func Not(arr *[5][10]int) {
	for i, v := range *arr {
        for j := range v {
			if ((*arr)[i][j] == 0) {
				(*arr)[i][j] = 1
			} else {
				(*arr)[i][j] = 0
			}
        }
    }
}

func main() {
    arr := [5][10]int{
        {1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
        {1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
        {0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        {1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
        {1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
    }
    Not(&arr)
    for _, v := range arr {
        fmt.Println(v)
    }
}