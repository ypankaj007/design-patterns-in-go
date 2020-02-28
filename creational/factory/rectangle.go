package factory

import "fmt"

// rectangle ..
//
type rectangle struct {
}

// Draw ..
func (rect *rectangle) Draw() {
	fmt.Println("Draw Rectangle")
}

// Fill ..
func (rect *rectangle) Fill(color string) {
	fmt.Println("Fill Rectangle with", color)
}

func GetRectangle() Shape {
	return &rectangle{}
}
