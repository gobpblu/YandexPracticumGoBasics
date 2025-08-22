package main

import "fmt"

func main() {

	a, b, c := false, false, false
	mainExp1 := !(a || b) || !c
	firstExp1 := (!a && !b) || !c
	secondExp1 := !a || !b || !c
	thirdExp1 := a || b && c
	fourthExp1 := (a && b) || c

	a, b, c = false, false, true
	mainExp2 := !(a || b) || !c
	firstExp2 := (!a && !b) || !c
	secondExp2 := !a || !b || !c
	thirdExp2 := a || b && c
	fourthExp2 := (a && b) || c

	a, b, c = false, true, false
	mainExp3 := !(a || b) || !c
	firstExp3 := (!a && !b) || !c
	secondExp3 := !a || !b || !c
	thirdExp3 := a || b && c
	fourthExp3 := (a && b) || c

	a, b, c = false, true, true
	mainExp4 := !(a || b) || !c
	firstExp4 := (!a && !b) || !c
	secondExp4 := !a || !b || !c
	thirdExp4 := a || b && c
	fourthExp4 := (a && b) || c

	a, b, c = true, false, false
	mainExp5 := !(a || b) || !c
	firstExp5 := (!a && !b) || !c
	secondExp5 := !a || !b || !c
	thirdExp5 := a || b && c
	fourthExp5 := (a && b) || c

	a, b, c = true, false, true
	mainExp6 := !(a || b) || !c
	firstExp6 := (!a && !b) || !c
	secondExp6 := !a || !b || !c
	thirdExp6 := a || b && c
	fourthExp6 := (a && b) || c

	a, b, c = true, true, false
	mainExp7 := !(a || b) || !c
	firstExp7 := (!a && !b) || !c
	secondExp7 := !a || !b || !c
	thirdExp7 := a || b && c
	fourthExp7 := (a && b) || c

	a, b, c = true, true, true
	mainExp8 := !(a || b) || !c
	firstExp8 := (!a && !b) || !c
	secondExp8 := !a || !b || !c
	thirdExp8 := a || b && c
	fourthExp8 := (a && b) || c

	isFirstExpEqual := mainExp1 == firstExp1 && mainExp2 == firstExp2 && mainExp3 == firstExp3 &&
		mainExp4 == firstExp4 && mainExp5 == firstExp5 && mainExp6 == firstExp6 && mainExp7 == firstExp7 &&
		mainExp8 == firstExp8

	isSecondExpEqual := mainExp1 == secondExp1 && mainExp2 == secondExp2 && mainExp3 == secondExp3 &&
		mainExp4 == secondExp4 && mainExp5 == secondExp5 && mainExp6 == secondExp6 && mainExp7 == secondExp7 &&
		mainExp8 == secondExp8

	isThirdExpEqual := mainExp1 == thirdExp1 && mainExp2 == thirdExp2 && mainExp3 == thirdExp3 &&
		mainExp4 == thirdExp4 && mainExp5 == thirdExp5 && mainExp6 == thirdExp6 && mainExp7 == thirdExp7 &&
		mainExp8 == thirdExp8

	isFourthExpEqual := mainExp1 == fourthExp1 && mainExp2 == fourthExp2 && mainExp3 == fourthExp3 &&
		mainExp4 == fourthExp4 && mainExp5 == fourthExp5 && mainExp6 == fourthExp6 && mainExp7 == fourthExp7 &&
		mainExp8 == fourthExp8

	a, b, c = true, true, true
	fmt.Printf("%t %t %t %t\n", isFirstExpEqual, isSecondExpEqual, isThirdExpEqual, isFourthExpEqual)
}

func checkExpressions(a, b, c bool) {
	expMain := !(a || b) || !c
	exp1 := (!a && !b) || !c
	exp2 := !a || !b || !c
	exp3 := a || b && c
	exp4 := (a && b) || c

	fmt.Printf("%t\t%t\t%t\t%t\n", expMain == exp1, expMain == exp2, expMain == exp3, expMain == exp4)
}
