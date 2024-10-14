package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(0)
	y := big.NewInt(0)

	x.Exp(big.NewInt(2), big.NewInt(30), nil)
	y.Exp(big.NewInt(2), big.NewInt(25), nil)

	fmt.Println(x)
	fmt.Println(y)

	res1 := big.NewInt(0)
	res1.Add(x, y)
	fmt.Println("x + y = ", res1)

	res2 := big.NewInt(0)
	res2.Sub(x, y)
	fmt.Println("x - y = ", res2)

	res3 := big.NewInt(0)
	res3.Mul(x, y)
	fmt.Println("x * y = ", res3)

	res4 := big.NewInt(0)
	res4.Div(x, y)
	fmt.Println("x / y = ", res4)

	res5 := big.NewInt(0)
	res5.Mod(x, y)
	fmt.Println("x % y = ", res5)
}
