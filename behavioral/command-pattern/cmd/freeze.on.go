package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type freezeON struct {
	freeze *devices.Freeze
}

// FreezeON //
func FreezeON(l *devices.Freeze) Command {
	return &freezeON{freeze: l}
}

func (freezeON *freezeON) Execute() {
	freezeON.freeze.On()
}
