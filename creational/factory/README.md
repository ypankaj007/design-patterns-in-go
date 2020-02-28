# Factory Design Pattern

We are going to create Shape interface and Rectangle, Square and Circle implements Shape. We have a ShapeFactory which has GetShape function. GetShape get a shape type as a param and return requested object hy type.
### UML Diagram
![alt text](https://raw.githubusercontent.com/ypankaj007/design-patterns-in-go/master/creational/factory/factory-design-pattern.png)

### Shape interface

```
type Shape interface {
    Draw()
    Fill(string)
}
```

### rectangle.go // implements shape inteface

```
import "fmt"

type rectangle struct {
}

func (rect *rectangle) Draw() {
	fmt.Println("Draw Rectangle")
}

func (rect *rectangle) Fill(color string) {
	fmt.Println("Fill Rectangle with", color)
}

func GetRectangle() Shape {
	return &rectangle{}
}
```
### square.go // implements shape inteface
```
import "fmt"

type square struct {
}

func (square *square) Draw() {
	fmt.Println("Draw Square")
}

func (square *square) Fill(color string) {
	fmt.Println("Fill Square with", color)
}

func GetSquare() Shape {
	return &square{}
}
```
### circle.go // implements shape inteface

```
import "fmt"

type circle struct {
}

func (circle *circle) Draw() {
	fmt.Println("Draw Circle")
}

func (circle *circle) Fill(color string) {
	fmt.Println("Fill Circle with", color)
}

func GetCircle() Shape {
	return &circle{}
}

```

### shape-factory.go // factory for shape

```
import "strings"

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
```

Now client can call GetShape function by passing required shape type object.
