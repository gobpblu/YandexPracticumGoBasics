package main

import "fmt"

func main() {
	fmt.Printf("%s\t%s\t%s\t%s\n", "a", "b", "c", "!(a || b) || !c")
	fmt.Println()

	a, b, c := false, false, false
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = false, false, true
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = false, true, false
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = true, false, false
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = false, true, true
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = true, false, true
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = true, true, false
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)

	a, b, c = true, true, true
	fmt.Printf("%t\t%t\t%t\t%t\n", a, b, c, !(a || b) || !c)
}
