package devices

import "fmt"

// Light ..
type Light struct{}

func (l *Light) On() {
	fmt.Println("LightOn command executed")
}

func (l *Light) Off() {
	fmt.Println("LightOff command executed")
}
