package main

import (
	"fmt"
)

const p = "ciao"

// constants multiople declaration

const (
	x = 10
	y = 20
	z = 30
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	const q = 11

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}
