package main

import "fmt"

func main() {

	s := struct {
		first     string
		num       map[string]int
		favDrinks []string
	}{
		first: "James",
		num: map[string]int{
			"home": 3103521123,
			"cell": 3109342211,
		},
		favDrinks: []string{
			"Vodka",
			"Coca Cola",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.num)
	fmt.Println(s.favDrinks)

	for k, v := range s.num {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
