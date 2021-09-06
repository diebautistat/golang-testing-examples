package calculator

func Add(a, b int) int {
	return a + b
}

type Calculator interface {
	Add(a, b int) int
}

type Calc struct{}

func (c *Calc)Add(a, b int) int {
	return a + b
}