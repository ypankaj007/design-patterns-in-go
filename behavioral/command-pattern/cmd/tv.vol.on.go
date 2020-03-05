package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type tvVolumeON struct {
	tv *devices.TV
}

// TvVolumeON //
func TvVolumeON(l *devices.TV) Command {
	return &tvVolumeON{tv: l}
}

func (tvVolumeON *tvVolumeON) Execute() {
	tvVolumeON.tv.VolumeOn()
}
