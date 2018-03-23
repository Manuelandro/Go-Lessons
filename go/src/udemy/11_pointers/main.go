package main

/* pointers ar type values which return the value at a specific address */

import (
	"fmt"
)

func main() {

	/* assign value */
	a := 55

	/* print value */
	fmt.Println(a)
	/* print memory address */
	fmt.Println(&a)

	/* assign the memory address to a pointer */
	var b *int = &a

	/* print the value of b */
	fmt.Println(b)
	/* print the value stored at the address */
	fmt.Println(*b)
}
