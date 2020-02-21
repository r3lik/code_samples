// https://play.golang.org/p/s03y19EOkUN

package main

import (
	"fmt"
)

var test_var = "mike"

const test_const = "mikey"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	
	test_var = "mike2"
	// Cannot change constant unlike var
	//test_const = "mikey2"
	fmt.Println(test_var)
	fmt.Println(test_const)
}

