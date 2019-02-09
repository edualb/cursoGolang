package main

import (
	"fmt"
	"reflect"
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
	*
	*					EXAMPLE 2
	*	Continuing...
	*
	*	name, course, module = "Nidel", "Docker Deep Dive", 3.2
	*
	*****************************************
	*
	*	name = "Nigel"
	*	course = "Docker Deep Dive"
	*	module = 3.2
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
	*
	*					EXAMPLE 2
	*
	*	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	*	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	*
	*****************************************
	*
	*					EXAMPLE 3
	*
	*	a := 10.000000 // float
	*	b := 3         //integer
	*
	*	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))
	*
	*	c := int(a) + b // 	We need cast the float number to calculate two integer numbers.
	*					//	But the A variable keep itself float.
	*
	*	fmt.Println("\n has value", c, "and is of type", reflect.TypeOf(c))
	 */

	a := 10.000000 // float
	b := 3         //integer

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := int(a) + b // We need cast the float number to calculate two integer numbers.

	fmt.Println("\n has value", c, "and is of type", reflect.TypeOf(c))
}
