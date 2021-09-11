package cooker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SmallStove struct {

}

func (s *SmallStove)Light() {
	fmt.Println("The stove is on! But on low fire.")
}

func (s *SmallStove)Cook(food string) string {
	fmt.Println("I'm cooking now! But slowly...")
	cookedFood := "cooked_" + food
	fmt.Println("I'm done cooking! It smells decent")
	return cookedFood
} 

func (s *SmallStove)TurnOff() {
	fmt.Println("The stove is off!")
}


func TestToCheckWhetherCookerKnowsHowToCookFish(t *testing.T) {
	food := "fish"
	stove := &SmallStove{}
	tested_cooker := &Cooker{Stove: stove}
	expected_food := "cooked_fish"
	cookerFood := tested_cooker.Cook(food)
	assert.Equal(t, expected_food, cookerFood)
}

func TestToCheckWhetherCookerKnowsHowToCookBeef(t *testing.T) {
	food := "beef"
	stove := &SmallStove{}
	tested_cooker := &Cooker{Stove: stove}
	expected_food := "cooked_beef"
	cookerFood := tested_cooker.Cook(food)
	assert.Equal(t, expected_food, cookerFood)
}