package main

import "fmt"

type person struct {
	name   string
	gender string
}

func main() {
	p1 := person{
		name:   "James Bond",
		gender: "male",
	}

	fmt.Println(p1)
	fmt.Printf("%p\n", &p1)
	// changes values at same address
	a1 := changeMe(&p1)
	fmt.Println(p1)
	fmt.Printf("%p\n", &p1)

	// print address where values were modified above
	fmt.Printf("%p\n", a1)

	// creates new address and changes values
	p2 := newAddress(p1)
	fmt.Println(p2)
	fmt.Printf("%p\n", &p2)
}

func changeMe(p *person) *person {
	(*p).name = "Ms Whoever"
	(*p).gender = "female"
	return p
}

func newAddress(p person) person {
	p.name = "Mr New"
	p.gender = "non binary"
	//fmt.Println(p)
	//fmt.Printf("%p\n", &p)
	return p
}
