package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	p3 := person{
		First: "Mr",
		Last:  "Rekt",
		Age:   35,
	}

	people := []person{
		p1,
		p2,
		p3,
	}

	fmt.Println(people)
	fmt.Println(people[(len(people) - 1)]) // print last item in slice

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
