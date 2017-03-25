package robot

import "github.com/mikhail-nikitin/go-robot/menu"

type CompositeCommand struct {
	commands []menu.Command
}

func (c *CompositeCommand) Execute() {
	for _, subCommand := range c.commands {
		subCommand.Execute()
	}
}

func (c *CompositeCommand) AppendCommand(command menu.Command) {
	c.commands = append(c.commands, command)
}