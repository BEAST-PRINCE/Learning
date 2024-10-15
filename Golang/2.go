package main

import "fmt"

func main() {
	// Arrays

	/*
		var arr [3]string
		fmt.Println(arr)

		arr2 := []string{"a", "b", "c", "d", "e"}
		fmt.Println(len(arr2), arr2)

		for i := 0; i < len(arr2); i++ {
			fmt.Println(arr2[i])
		}

		var s []string = arr2[1:4]
		fmt.Println(len(s), s)
		fmt.Println(s[:cap(s)])

		var arr3 []int = []int{1, 2, 3, 4}
		arr3 = append(arr3, 5)
		fmt.Println(len(arr3), arr3)

		arr4 := make([]int, 8) // Make created a Slice
		fmt.Printf("%T\n", arr4)
		fmt.Println(len(arr4), arr4)

		for i, element := range arr2 {
			fmt.Println(i, element)
		}
	*/

	//////////////////////////////////////////////////////////

	// Maps

	var mp map[string]int = map[string]int{"z": 26, "a": 1, "b": 2, "c": 3, "d": 4} // MAps are unordered
	fmt.Println(len(mp), mp)
	fmt.Println(mp["z"])
	mp["z"] = -1
	fmt.Println(mp["z"])
	delete(mp, "z")
	delete(mp, "p")
	fmt.Println(mp)

	mp1 := make(map[string]int)
	fmt.Println(len(mp1), mp1)
	mp1["PC"] = 500
	fmt.Println(len(mp1), mp1)

	val, ok := mp1["Laptop"]
	fmt.Println(val, ok)

}
