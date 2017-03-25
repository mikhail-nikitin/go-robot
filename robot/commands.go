package robot

type TurnOnCommand struct {
	Robot *Robot
}

func (c *TurnOnCommand) Execute() {
	c.Robot.TurnOn()
}

type TurnOffCommand struct {
	Robot *Robot
}

func (c *TurnOffCommand) Execute() {
	c.Robot.TurnOff()
}

type MoveSpecificDirectionCommand struct {
	Robot *Robot
	Direction int64
}

func (c *MoveSpecificDirectionCommand) Execute() {
	c.Robot.Walk(c.Direction)
}

type StatusCommand struct {
	Robot *Robot
}

func (c *StatusCommand) Execute() {
	c.Robot.Status()
}

type StopCommand struct {
	Robot *Robot
}

func (c *StopCommand) Execute() {
	c.Robot.Stop()
}