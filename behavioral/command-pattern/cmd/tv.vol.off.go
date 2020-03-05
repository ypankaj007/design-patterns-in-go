package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type tvVolumeOff struct {
	tv *devices.TV
}

// TvVolumeOff //
func TvVolumeOff(l *devices.TV) Command {
	return &tvVolumeOff{tv: l}
}

func (tvVolumeOff *tvVolumeOff) Execute() {
	tvVolumeOff.tv.VolumeOff()
}
