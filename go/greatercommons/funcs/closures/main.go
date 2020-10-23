package main

import "fmt"

func runIncrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func incrementor() {
	a := runIncrementor()
	b := runIncrementor()
	fmt.Printf("%T\n", a)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func main() {
	var x int // limit scope to function
	fmt.Println(x)
	x++

	// limit scope in code-block
	{
		y := 42
		fmt.Println(y)
	}

	// fmt.Println(y) <-- will error as out of scope

	fmt.Println(x)
	incrementor()
}
