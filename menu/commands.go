package menu

type ShowHelpCommand struct {
	Menu *Menu
}

func (c *ShowHelpCommand) Execute() {
	c.Menu.ShowInstructions()
}

type ExitMenuCommand struct {
	Menu *Menu
}

func (c *ExitMenuCommand) Execute() {
	c.Menu.Exit()
}
