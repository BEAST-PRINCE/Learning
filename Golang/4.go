package main

import "fmt"

func changeVal(s1 *string, s2 string) {
	*s1 = "New Value"
	s2 = "New Value"
}

func main() {
	// Pointers and Dereference
	fmt.Println("Starting")

	x := 7
	fmt.Println(&x)

	y := &x // Pointer
	fmt.Println(&y)

	*y = 10 // Dereference
	fmt.Println(x, y)

	val1 := "String"
	val2 := "String"
	changeVal(&val1, val2)
	fmt.Println(val1, val2)

	val := "Namaste"
	var ptr *string = &val // Pointer
	fmt.Println(ptr, val)
	fmt.Println(*ptr, val)
	fmt.Println(ptr, &ptr)
}
