package main

import (
	"fmt"
	"golangtestingexamples/cooker"
)

func main(){
	stove := &cooker.Stove{}
	cooker := &cooker.Cooker{Stove: stove}
	food := "chicken"
	cookedFood := cooker.Cook(food)
	fmt.Println(cookedFood)
}