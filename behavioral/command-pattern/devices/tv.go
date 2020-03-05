package devices

import "fmt"

// TV ..
type TV struct{}

func (f *TV) On() {
	fmt.Println("TVOn command executed")
}

func (f *TV) Off() {
	fmt.Println("TVOff command executed")
}

func (f *TV) VolumeOn() {
	fmt.Println("TV VolumeOn command executed")
}

func (f *TV) VolumeOff() {
	fmt.Println("TV VolumeOff command executed")
}
