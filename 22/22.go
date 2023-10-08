package main

import (
	"fmt"
	"math/big"
)

/* Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20. */

func main() {
	var a, b int64 = 1048577214124124, 10485785126534
	fmt.Println("Mul:", mul(a, b))
	fmt.Println("Div:", div(a, b))
	fmt.Println("Add:", add(a, b))
	fmt.Println("Sub:", sub(a, b))
}

func mul(a, b int64) big.Int {
	var c big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	c.Mul(a1, b1)
	return c
}

func div(a, b int64) big.Int {
	var c big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	c.Div(a1, b1)
	return c
}

func add(a, b int64) big.Int {
	var c big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	c.Add(a1, b1)
	return c
}

func sub(a, b int64) big.Int {
	var c big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	c.Sub(a1, b1)
	return c
}
