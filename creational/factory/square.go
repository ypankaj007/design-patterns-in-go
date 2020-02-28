package main

import "fmt"

// square ..
//
type square struct {
}

// Draw ..
func (square *square) Draw() {
	fmt.Println("Draw Square")
}

// Fill ..
func (square *square) Fill(color string) {
	fmt.Println("Fill Square with", color)
}

func GetSquare() Shape {
	return &square{}
}
