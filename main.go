package main

import (
	"bufio"
	"os"
	"github.com/mikhail-nikitin/go-robot/menu"
	"github.com/mikhail-nikitin/go-robot/robot"
)

func main() {
	r := robot.New()
	m := menu.New()

	m.AddItem("on", "Turns the Robot on", &robot.TurnOnCommand{r})
	m.AddItem("off", "Turns the Robot off", &robot.TurnOffCommand{r})

	m.AddItem("up", "Makes the Robot walk up", &robot.MoveSpecificDirectionCommand{r, robot.UP})
	m.AddItem("down", "Makes the Robot walk down", &robot.MoveSpecificDirectionCommand{r, robot.DOWN})
	m.AddItem("left", "Makes the Robot walk left", &robot.MoveSpecificDirectionCommand{r,robot.LEFT})
	m.AddItem("right", "Makes the Robot walk right", &robot.MoveSpecificDirectionCommand{r, robot.RIGHT})

	// m.AddItem("horse_moving", "Makes the Robot walk like horse", ???)

	m.AddItem("status", "Prints Robot status (turned on/off, walk direction)", &robot.StatusCommand{r})
	m.AddItem("stop", "Stops the Robot", &robot.StopCommand{r})

	helpCommand := &menu.ShowHelpCommand{m}
	m.AddItem("help", "Show instructions", helpCommand)
	m.AddItem("exit", "Exit from this m", &menu.ExitMenuCommand{m})

	helpCommand.Execute()
	m.Run(bufio.NewReader(os.Stdin))
}
