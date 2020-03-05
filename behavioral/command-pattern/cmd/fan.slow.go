package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type fanSlow struct {
	fan *devices.Fan
}

// FanSlow //
func FanSlow(l *devices.Fan) Command {
	return &fanSlow{fan: l}
}

func (fanSlow *fanSlow) Execute() {
	fanSlow.fan.Slow()
}
