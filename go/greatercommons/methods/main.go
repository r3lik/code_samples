package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (returns(s)) <code>

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, s.ltk)

}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		// can also use short form and exlude fieldnames
		//person: person{
		//			"James",
		//			"Bond",
		//		}
		ltk: true,
	}
	//fmt.Println(sa1)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.ltk)
	sa1.speak()
}
