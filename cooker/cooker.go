package cooker

type Cooker struct {
	Stove *Stove
}

func (c *Cooker) Cook(food string) (string) {
	c.Stove.Light()
	cookedFood := c.Stove.Cook(food)
	c.Stove.TurnOff()
	return cookedFood
}