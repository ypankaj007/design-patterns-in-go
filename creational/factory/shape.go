package factory

// Shape ..
// shape interface has a common
// func which all shapes has "draw"
type Shape interface {
	Draw()
	Fill(string)
}
