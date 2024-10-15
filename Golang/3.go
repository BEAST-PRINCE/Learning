package main

import (
	"fmt"
)

func f1(x string) {
	fmt.Println("Golang Functions")
	fmt.Println("Welcome ", x)
}

func add(a int, b int) int {
	return a + b
}

func f2(a int, b int) (int, int) {
	defer fmt.Println("Result generated") // defer delays the execution of this line until the function returns
	fmt.Println("Generating result...")
	return a + b, a * b
}

func f3(a int, b int) (r1 int, r2 int) {
	r1 = a + b
	r2 = a - b
	return
}

func f4(myfunc func() string) { // Passing a function as parameter to another function
	myfunc()
}

func func_closure(x string) func() { // This is the closure for a function
	return func() { fmt.Println(x) }
}

func main() {
	f1("BEAST PRINCE")

	result := add(5, 15)
	fmt.Println(result)

	r1, r2 := f2(5, 10)
	fmt.Println(r1, r2)

	r1, r2 = f3(3, 5)
	fmt.Println(r1, r2)

	fun_reference := f2 // referencing a function
	fun_reference(10, 20)

	f := func() string {
		fmt.Println("Function F")
		return "Val"
	}

	_ = f() // It is same as calling the function directly without assigning the returned value to a variable

	x := func(y int) int {
		return y * y
	}(5) // directly calling the function and storing the result in a variable
	fmt.Println(x)

	f4(f) // Passing function as argument to another function

	func_closure("Closure Function Called")() // We need to call the returned function again in order to get the output
	temp := func_closure("Closure Function Called Again")
	temp()
}
