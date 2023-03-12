package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  // 42
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int

	var b *int = &a
	fmt.Println(b)   // & gives you the address
	fmt.Println(*b)  // * gives your value stored at an address, 42
	fmt.Println(*&a) // 42

	*b = 43 // derefence address of a and set that value to 43
	fmt.Println(a)
}
