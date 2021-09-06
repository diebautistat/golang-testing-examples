package student

import (
	"testing"

	//"github.com/stretchr/testify/assert"
)

type MockCalculator struct {
	addMethodCalled bool
	firstArg int
	secondArg int
}

func (mc *MockCalculator)Add(a, b int) int {
	mc.firstArg = a
	mc.secondArg = b
	mc.addMethodCalled = true
	return 0
}


func TestThatStudentKnowsHowToAddUsingACalculator(t *testing.T){
	a := 3
	b := 4
	calculator := MockCalculator{}
	studnt := Student{calc: &calculator}
	_ = studnt.Add(a, b)
	if calculator.addMethodCalled != true{
		t.Fail()
	}
	
}

func TestThatStudentKnowsHowToAddUsingACalculatorWithCorrectInputs(t *testing.T){
	a := 3
	b := 4
	calculator := MockCalculator{}
	studnt := Student{calc: &calculator}
	_ = studnt.Add(a, b)
	if calculator.firstArg != a {
		t.Fail()
	}
	if calculator.secondArg != b {
		t.Fail()
	}	
}