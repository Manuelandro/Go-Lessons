package main

import "fmt"

// classic variables - var and initialization
var name = "Manuel"
var lastname string
var age int32 = 20
var male bool

func main() {

	// shortand variables - only works inside funcs
	a := 1
	b := true
	c := "string"
	d := 6.8

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
}
