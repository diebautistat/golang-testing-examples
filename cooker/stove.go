package cooker

import "fmt"

type CookingStove interface {
	Light()
	TurnOff()
	Cook(string) string
}

type Stove struct {
	lightBulb bool
}

func (s *Stove)Light() {
	fmt.Println("The stove is on! Be careful!")
	s.turnLightBulbOn()
}

func (s *Stove)Cook(food string) string {
	fmt.Println("I'm cooking now!")
	cookedFood := "cooked_" + food
	fmt.Println("I'm done cooking! It smells really good!")
	return cookedFood
} 

func (s *Stove)TurnOff() {
	s.turnLightBulbOff()
	fmt.Println("The stove is off! You can rest now.")
}

func (s *Stove)turnLightBulbOn() {
	s.lightBulb = true
	fmt.Println("The stove light bulb is on! Enjoy!")
}

func (s *Stove)turnLightBulbOff() {
	s.lightBulb = false
	fmt.Println("The stove light bulb is off! Good bye!")
}