package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type tvON struct {
	tv *devices.TV
}

// TvON //
func TvON(l *devices.TV) Command {
	return &tvON{tv: l}
}

func (tvON *tvON) Execute() {
	tvON.tv.On()
}
