package main

import (
	"design-patterns-in-go/behavioral/command-pattern/cmd"
	"design-patterns-in-go/behavioral/command-pattern/devices"
	"design-patterns-in-go/behavioral/command-pattern/remotes"
	"fmt"
)

func main() {
	remote := remotes.GetRemote()
	light := &devices.Light{}
	lightOn := cmd.LightON(light)
	remote.SetCommand(lightOn)
	lightOff := cmd.LightOff(light)
	remote.SetCommand(lightOff)
	fan := &devices.Fan{}
	fanOn := cmd.FanON(fan)
	remote.SetCommand(fanOn)
	fanOff := cmd.FanOff(fan)
	remote.SetCommand(fanOff)
	fanSlow := cmd.FanSlow(fan)
	remote.SetCommand(fanSlow)
	freeze := &devices.Freeze{}
	freezeOn := cmd.FreezeON(freeze)
	remote.SetCommand(freezeOn)
	freezeOff := cmd.FreezeOff(freeze)
	remote.SetCommand(freezeOff)
	tv := &devices.TV{}
	tvOn := cmd.TvON(tv)
	remote.SetCommand(tvOn)
	tvOff := cmd.TvOff(tv)
	remote.SetCommand(tvOff)
	tvVolumeOn := cmd.TvVolumeON(tv)
	remote.SetCommand(tvVolumeOn)
	tvVolumeOff := cmd.TvVolumeOff(tv)
	remote.SetCommand(tvVolumeOff)
	fmt.Println("Start executing commands")
	remote.ExecuteCommand()
}
