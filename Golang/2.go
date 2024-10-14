package main

import (
	"fmt"
)

func main() {
	// Arrays

	// var arr [3]string
	// fmt.Println(arr)

	arr2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(arr2), arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
}
