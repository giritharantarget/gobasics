package main

import "fmt"

func main() {
	fmt.Println("Hello World")


	// Primitive data type
	r := 42
	fmt.Println(r)

	//  Pointer operator

	firstName := "Hello"
	reference := & firstName

	fmt.Println("reference" , reference)

	// Constants

	const a int  = 3;
	fmt.Println(a +3)

	//goCollectionsArrays()

	Slice();
	MapType();
	Struct();
}
