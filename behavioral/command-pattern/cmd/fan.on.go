package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type fanON struct {
	fan *devices.Fan
}

// FanON //
func FanON(l *devices.Fan) Command {
	return &fanON{fan: l}
}

func (fanON *fanON) Execute() {
	fanON.fan.On()
}
