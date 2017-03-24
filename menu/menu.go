package menu

import (
	"fmt"
	"bufio"
	"io"
)

type Command interface {
	Execute()
}

type Menu struct {
	exit  bool
	items map[string]Item
}

type Item struct {
	shortcut    string
	description string
	command     Command
}

func (m *Menu) AddItem(shortcut, description string, command Command) {
	m.items[shortcut] = Item{shortcut, description, command}
}

func (m *Menu) Run(input *bufio.Reader) {
	for {
		s, isPrefix, err := input.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			fmt.Println("Command is too long, try again")
			continue
		}
		if !m.executeCommand(string(s)) {
			break
		}
	}
}

func (m *Menu) executeCommand(word string) bool {
	m.exit = false
	item, ok := m.items[word]
	if !ok {
		fmt.Println("Unknown command")
	} else {
		item.command.Execute()
	}
	return !m.exit
}

func (m *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range m.items {
		fmt.Printf("\t%v: %v\n", item.shortcut, item.description)
	}
}

func New() *Menu {
	m := &Menu{}
	m.items = make(map[string]Item)
	return m
}

func (m *Menu) Exit() {
	m.exit = true
}