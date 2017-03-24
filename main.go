package main

import (
	"bufio"
	"os"
	"github.com/mikhail-nikitin/go-robot/menu"
)

func main() {
	// robot := NewRobot()
	m := menu.New()

	// TODO: implement all or some commands

	// m.AddItem("on", "Turns the Robot on", ???)
	// m.AddItem("off", "Turns the Robot off", ???)

	// m.AddItem("up", "Makes the Robot walk up", ???)
	// m.AddItem("down", "Makes the Robot walk down", ???)
	// m.AddItem("left", "Makes the Robot walk left", ???)
	// m.AddItem("right", "Makes the Robot walk right", ???)

	// m.AddItem("horse_moving", "Makes the Robot walk like horse", ???)

	// m.AddItem("status", "Prints Robot status (turned on/off, walk direction)", ???)
	// m.AddItem("stop", "Stops the Robot", ???)

	helpCommand := &menu.ShowHelpCommand{m}
	m.AddItem("help", "Show instructions", helpCommand)
	m.AddItem("exit", "Exit from this m", &menu.ExitMenuCommand{m})

	helpCommand.Execute()
	m.Run(bufio.NewReader(os.Stdin))
}
