package main

import "fmt"

type person struct {
	first    string
	last     string
	favFlave []string
}

func main() {
	p1 := person{
		first:    "Mike",
		last:     "Black",
		favFlave: []string{"Coffee", "Lemon"},
	}
	p2 := person{
		first:    "Marie",
		last:     "Black",
		favFlave: []string{"Vanilla", "Chocolate"},
	}
	// more complicated way of creating a slice of person and grouping the persons
	people := []person{p1, p2}
	for _, p := range people {
		fmt.Printf("First: %s, Last: %s, Favorites: %v\n", p.first, p.last, p.favFlave)
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.favFlave {
		fmt.Println(v)
	}
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.favFlave {
		fmt.Println(v)
	}
}
