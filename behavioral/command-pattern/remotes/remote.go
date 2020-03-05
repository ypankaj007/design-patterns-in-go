package remotes

import "design-patterns-in-go/behavioral/command-pattern/cmd"

// Remote ..
type Remote interface {
	SetCommand(cmd.Command)
	ExecuteCommand()
}

type remote struct {
	commands []cmd.Command
}

// GetRemote ..
func GetRemote() Remote {
	return &remote{}
}

func (r *remote) SetCommand(c cmd.Command) {
	r.commands = append(r.commands, c)
}

func (r *remote) ExecuteCommand() {
	for _, command := range r.commands {
		command.Execute()
	}
	r.commands = nil
}
