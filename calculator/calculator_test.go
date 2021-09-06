package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
	

func TestBasicAddingCase(t *testing.T){
	//Arrange
	a := 5
	b := 10
	expected_result := 15
	//Act
	result := Add(a, b)
	//Assert
	if expected_result != result {
		t.Fail()
	}
	//assert.Equal(t, expected_result, result)
}

func TestAddingUsingAStruct(t *testing.T){
	//Arrange
	a := 5
	b := 10
	expected_result := 15
	calc := Calc{}
	//Act
	result := calc.Add(a, b)
	//Assert
	assert.Equal(t, expected_result, result)
}