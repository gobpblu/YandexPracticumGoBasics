package main

import "fmt"

const (
	CmdAdd   = iota // сложить два числа в стеке (top-1) = (top-1) + top
	CmdSub          // вычесть (top-1) = (top-1) - top
	CmdMul          // умножить (top-1) = (top-1) * top
	CmdDiv          // разделить (top-1) = (top-1) / top
	CmdPush         // поместить следующее число в стек
	CmdPop          // убрать число из стека
	CmdPrint        // печать верхнего элемента стека
	CmdSave         // сохранить число из стека в ячейку
	CmdLoad         // переместить число из ячейки в стек
)

func main() {
	program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush,
		-13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul,
		CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7,
		CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
		CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub,
		CmdPush, 1000, CmdDiv, CmdPrint}
	registers := make(map[rune]int)
	stack := make([]int, 0, 100)

	for i := 0; i < len(program); i++ {
		cmd := program[i]
		switch cmd {
		case CmdAdd:
			if len(stack) < 2 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			top := len(stack) - 1
			stack[top-1] += stack[top]
			stack = stack[:top]

		case CmdSub:
			if len(stack) < 2 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			top := len(stack) - 1
			stack[top-1] -= stack[top]
			stack = stack[:top]

		case CmdMul:
			if len(stack) < 2 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			top := len(stack) - 1
			stack[top-1] *= stack[top]
			stack = stack[:top]

		case CmdDiv:
			if len(stack) < 2 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			top := len(stack) - 1
			if stack[top] == 0 {
				fmt.Println("На ноль делить нельзя!!")
				return
			}
			stack[top-1] /= stack[top]
			stack = stack[:top]

		case CmdPush:
			if len(program) > i+1 {
				stack = append(stack, program[i+1])
				i++
			}

		case CmdPop:
			if len(stack) < 1 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			stack = stack[:len(stack)-1]

		case CmdPrint:
			if len(stack) < 1 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			fmt.Println(stack[len(stack)-1])

		case CmdSave:
			if len(stack) < 1 {
				fmt.Println("Что-то пошло не так!")
				return
			}
			if len(program) > i+1 {
				key := rune(program[i+1])
				topValue := stack[len(stack)-1]
				registers[key] = topValue
			}

		case CmdLoad:
			if len(program) > i+1 {
				key := rune(program[i+1])
				value, ok := registers[key]
				if !ok {
					fmt.Println("Что-то пошло не так!")
					return
				}
				stack = append(stack, value)
			}
		}
	}
}
