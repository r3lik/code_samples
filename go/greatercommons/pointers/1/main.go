package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * gives your value stored at an address
	fmt.Println(*&a)

	*b = 43 // derefence address of a and set that value to 43
	fmt.Println(a)
}
