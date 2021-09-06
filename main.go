package main

import (
	"fmt"
	"golangtestingexamples/calculator"
)

func main(){
	a := 1
	b := 2
	sum := calculator.Add(a, b)
	fmt.Printf("The value of the sum is %d", sum)
}