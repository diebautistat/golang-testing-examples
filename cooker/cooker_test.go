package cooker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCheckWhetherCookerKnowsHowToCookFish(t *testing.T) {
	food := "fish"
	stove := &Stove{}
	tested_cooker := &Cooker{Stove: stove}
	expected_food := "cooked_fish"
	cookerFood := tested_cooker.Cook(food)
	assert.Equal(t, expected_food, cookerFood)
}

func TestToCheckWhetherCookerKnowsHowToCookBeef(t *testing.T) {
	food := "beef"
	stove := &Stove{}
	tested_cooker := &Cooker{Stove: stove}
	expected_food := "cooked_beef"
	cookerFood := tested_cooker.Cook(food)
	assert.Equal(t, expected_food, cookerFood)
}