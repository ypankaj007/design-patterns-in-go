package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type tvOff struct {
	tv *devices.TV
}

// TvOff //
func TvOff(l *devices.TV) Command {
	return &tvOff{tv: l}
}

func (tvOff *tvOff) Execute() {
	tvOff.tv.Off()
}
