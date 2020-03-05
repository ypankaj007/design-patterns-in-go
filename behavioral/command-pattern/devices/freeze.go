package devices

import "fmt"

// Freeze ..
type Freeze struct{}

func (f *Freeze) On() {
	fmt.Println("FreezeOn command executed")
}

func (f *Freeze) Off() {
	fmt.Println("FreezeOff command executed")
}
