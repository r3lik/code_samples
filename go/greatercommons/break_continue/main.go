// https://play.golang.org/p/etkrZ8fpYWr
package main

import (
	"fmt"
)

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		// continue to next iteration of loop and skip code in between (Print) if number not even
		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("done.")
}
