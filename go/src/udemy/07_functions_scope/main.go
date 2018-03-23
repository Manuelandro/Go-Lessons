package main

import (
	"fmt"
)

func sum() int {
	x := 10
	increment := func() int {
		return x + 10
	}

	return increment()
}

func initialize() func() int {
	x := 10
	return func() int {
		x++
		return x
	}
}

func main() {

	summa := sum()
	initialized1 := initialize()
	initialized2 := initialize()
	fmt.Println(summa)
	fmt.Println(initialized1())
	fmt.Println(initialized2())
}
