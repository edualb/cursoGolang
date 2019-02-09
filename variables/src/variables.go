package main

import "fmt"

/*					INFOS:
*	-	Package level variables are global.
*	-	Shorthand declare and initializa ":=" only workd inside of functions
*	-	Variables declared at the function level MUST be used (Garbage Collector)
 */

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
)

func main() {
	/*
	*					EXAMPLE 1
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
	*
	*****************************************
	*
	*					EXAMPLE 4
	*
	*	name := "Nigel"
	*	course := "Docker Deep Dive"
	*	module := 3.2
	*
	*	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	*	fmt.Println("Course is set to", course, "and is of type", reflect.TypeOf(course))
	*	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	*
	*****************************************
	*
	*					EXAMPLE 5
	*	Pointers:
	*
	*	module := 3.2
	*	ptr := &module // The "&" is a memory reference of module variable
	*
	*	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	*
	*	// ptr = memory reference
	*	// *ptr = memory dereferencing
	*	fmt.Println("Memory address of *module* variable is", ptr,
	*		"and the value of *module* is", *ptr)
	*
	*****************************************
	*					EXAMPLE 6
	*
	*	name := "Nigel"
	*	course := "Docker Deep Dive"
	*
	*	fmt.Println("\nHi", name, "you're currently watching",
	*	course)
	*
	*	// we create this function to show how to work with Pointers
	*	changeCourse(&course)
	*
	*	fmt.Println("\nYou are now watching course", course)
	 */

	name := "Nigel"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching",
		course)

	// we create this function to show how to work with Pointers
	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

/*
*						EXAMPLE 6
*
*	This functions is used to show how to work with pointers
*
*	func changeCourse(courseParam *string) string {
*	*courseParam = "First Look: Native Docker Clustering"
*
*	fmt.Println("\nTrying to change your course to", *courseParam)
*
*	return *courseParam
*	}
 */

func changeCourse(courseParam *string) string {
	*courseParam = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *courseParam)

	return *courseParam
}
