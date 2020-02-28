package factory

import "fmt"

// circle ..
//
type circle struct {
}

// Draw ..
func (circle *circle) Draw() {
	fmt.Println("Draw Circle")
}

// Fill ..
func (circle *circle) Fill(color string) {
	fmt.Println("Fill Circle with", color)
}

func GetCircle() Shape {
	return &circle{}
}
