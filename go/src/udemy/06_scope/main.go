// LEVELS OF SCOPOE

// • universe level
// • package level
// • file level
// • block level

package main

import (
	"fmt" // file level scope
)

var x int // package level scope

func main() {

	y := 10 // block level scope
	fmt.Printf("%d", y)
}
