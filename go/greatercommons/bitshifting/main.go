// https://play.golang.org/p/sVmr6r7-WvB
package main

import (
	"fmt"
)

// bitshifting by 10, 20, 30 etc with iota (0,1,2..)
// a bit too clever and hurts readbility
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
