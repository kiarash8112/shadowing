package t

import "fmt"

func noShadows() {
	if true {
		x := 20
		fmt.Println(x)
	}
	if true {
		x := 20
		fmt.Println(x)
	}
}

func noShadows1() {
	y := 1
	if true {
		x := 20
		fmt.Println(x)
	}
	fmt.Println(y)
}

func Shadows() {
	x := 10

	if true {
		x := 20 // want `variable "x" shadows an existing variable`
		fmt.Println(x)
	}

	fmt.Println(x)
}

func Shadows1() {
	x := 10
	if true {
		if true {
			y := 10
			fmt.Println(y)
		}
		if true {
			x := 20 // want `variable "x" shadows an existing variable`
			fmt.Println(x)
		}
	}

	fmt.Println(x)
}
