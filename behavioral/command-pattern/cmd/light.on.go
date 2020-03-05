package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type lightON struct {
	light *devices.Light
}

// LightON //
func LightON(l *devices.Light) Command {
	return &lightON{light: l}
}

func (lightON *lightON) Execute() {
	lightON.light.On()
}
