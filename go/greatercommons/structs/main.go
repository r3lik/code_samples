package main

import "fmt"

// structs let us aggregate values of different types
// composite data structure
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   26,
	}

	// embedded struct
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	// anonymous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(sa1)
	fmt.Println(p1)
	fmt.Println("Anonymous:", p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
