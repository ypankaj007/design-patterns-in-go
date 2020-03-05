package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type lightOff struct {
	light *devices.Light
}

// LightOff //
func LightOff(l *devices.Light) Command {
	return &lightOff{light: l}
}

func (lightOff *lightOff) Execute() {
	lightOff.light.Off()
}
