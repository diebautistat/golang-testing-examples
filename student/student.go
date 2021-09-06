package student

import "golangtestingexamples/calculator"

type Student struct {
	calc calculator.Calculator
}

func (s *Student)Add(a, b int) int {
	result := s.calc.Add(a, b)
	return result
}