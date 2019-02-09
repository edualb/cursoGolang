package main

import (
	"fmt"
)

var (
	/*					EXAMPLE 1
	*	You can declare package level variables like this:
	*
	*	name string
	*	course string
	*	module float64
	*
	*****************************************
	*
	*	name, course string
	*	module float64
	*
	*****************************************
	 */
	name   string
	course string
	module float64
)

func main() {
	/*					EXAMPLE 1
	*	If you try run the code below, the name will be a void string
	*	and the module will be the number 0
	*
	*	fmt.Println("Name is set to", name)
	*	fmt.Println("Module is set to", module)
	*
	*****************************************
	 */
	fmt.Println("Name is set to", name)
	fmt.Println("Module is set to", module)
}
