package factory

import "strings"

// GetShape is a factory method
// thats return shape object depend on shape type
func GetShape(shapeType string) Shape {
	switch strings.ToLower(shapeType) {
	case "rectangle":
		return GetRectangle()
	case "square":
		return GetSquare()
	case "circle":
		return GetCircle()
	default:
		return nil
	}
}
