package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		luxury: false,
	}
	fmt.Println(t1, "Color:", t1.color)
	fmt.Println(s1, "Color:", s1.color)
}
