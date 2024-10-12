package main

import "fmt"

func main() {
	fmt.Println("Namaste India!!")
	fmt.Print("Learning Go :)\n")

	// var <name> type?
	var name string = "Beast_Prince"
	var alphabet = 'A' // characters by default are stores int of their unicode equivalents
	alphabet2 := 'a'
	fmt.Println("I'm ", name)
	fmt.Printf("%T  %T\n", alphabet, alphabet2) // %T gives type of the variable
	fmt.Println(alphabet, alphabet2)

	var x = fmt.Sprintf("%T  %T  %v\n", alphabet, alphabet2, alphabet2)
	fmt.Println(x)
	// %t is to print boolean values
	// %b base2(binary)  %o base8(octal)  %d base10(decimal)  %x/%X base16(hexadecimal)
	// %s default string    %q for strings with double quotes

}
