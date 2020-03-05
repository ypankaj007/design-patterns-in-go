package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type fanOff struct {
	fan *devices.Fan
}

// FanOff //
func FanOff(l *devices.Fan) Command {
	return &fanOff{fan: l}
}

func (fanOff *fanOff) Execute() {
	fanOff.fan.Off()
}
