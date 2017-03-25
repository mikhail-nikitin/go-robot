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