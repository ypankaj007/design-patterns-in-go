package main

import (
	"design-patterns-in-go/creational/factory"
)

func main() {

	// ****************  Factory Design Patten - Start **********************
	circle := factory.GetShape("circle")
	circle.Draw()
	circle.Fill("red")

	rectangle := factory.GetShape("rectangle")
	rectangle.Draw()
	rectangle.Fill("blue")

	square := factory.GetShape("square")
	square.Draw()
	square.Fill("green")

	// ****************  Factory Design Patten - End **********************
}
