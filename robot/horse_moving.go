package robot

import "github.com/mikhail-nikitin/go-robot/menu"

func NewHorseMovingCommand(robot *Robot) (result menu.Command) {
	c := &CompositeCommand{}
	c.AppendCommand(&MoveSpecificDirectionCommand{robot, UP})
	c.AppendCommand(&MoveSpecificDirectionCommand{robot, RIGHT})
	c.AppendCommand(&MoveSpecificDirectionCommand{robot, RIGHT})
	return c
}