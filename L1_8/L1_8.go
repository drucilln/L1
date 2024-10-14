package main

import (
	"flag"
	"fmt"
)

func main() {
	var x uint64
	x = 8
	bitPos := flag.Int("b", 4, "bit position")
	flag.Parse()
	fmt.Printf("x=%064b\n", x)
	x ^= 1 << *bitPos
	fmt.Printf("x=%064b\n", x)
}
