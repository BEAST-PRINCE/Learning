package main

import "fmt"

// Struct
type Id struct {
	n    int
	name string
}

func changeName(i *Id, s string) {
	i.name = s
}

// Struct Method
func (i Id) getName() string {
	return i.name
}

func main() {
	var id1 Id = Id{1, "Abhi"}

	/*
		fmt.Println(id1)
		fmt.Printf("%T\n", id1)
		fmt.Println(id1.n, id1.name)
	*/
	ptr := &id1
	// fmt.Println(ptr, *ptr)

	changeName(ptr, "Abhisek")
	fmt.Println(id1, ptr)

	fmt.Println(id1.getName())
}
