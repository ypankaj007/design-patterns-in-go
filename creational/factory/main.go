package main

func main() {

	// ****************  Factory Design Patten - Start **********************
	circle := GetShape("circle")
	circle.Draw()
	circle.Fill("red")

	rectangle := GetShape("rectangle")
	rectangle.Draw()
	rectangle.Fill("blue")

	square := GetShape("square")
	square.Draw()
	square.Fill("green")

	// ****************  Factory Design Patten - End **********************
}
