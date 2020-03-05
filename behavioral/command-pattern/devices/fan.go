package devices

import "fmt"

// Fan ..
type Fan struct{}

func (f *Fan) On() {
	fmt.Println("FanOn command executed")
}

func (f *Fan) Off() {
	fmt.Println("FanOff command executed")
}

func (f *Fan) Slow() {
	fmt.Println("FanSlow command executed")
}
