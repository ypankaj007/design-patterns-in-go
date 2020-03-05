package cmd

import "design-patterns-in-go/behavioral/command-pattern/devices"

type freezeOff struct {
	freeze *devices.Freeze
}

// FreezeOff //
func FreezeOff(l *devices.Freeze) Command {
	return &freezeOff{freeze: l}
}

func (freezeOff *freezeOff) Execute() {
	freezeOff.freeze.Off()
}
