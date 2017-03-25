package robot

import "fmt"

const (
	UP int64 = iota
	DOWN
	LEFT
	RIGHT
	NO_DIRECTION
)

type Robot struct {
	direction int64
	turnedOn  bool
}

func (r *Robot) TurnOn() {
	if !r.turnedOn {
		r.turnedOn = true
		fmt.Println("It am waiting for your commands")
	}
}

func (r *Robot) TurnOff() {
	if r.turnedOn {
		r.turnedOn = false
		r.direction = NO_DIRECTION
		fmt.Println("It is a pleasure to serve you")
	}
}

func (r *Robot) Walk(direction int64) {
	if r.turnedOn {
		r.direction = direction

		fmt.Printf("Walking %v\n", directionToString(direction))
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Stop() {
	if r.turnedOn {
		if r.direction != NO_DIRECTION {
			r.direction = NO_DIRECTION
			fmt.Printf("Stopped\n")
		} else {
			fmt.Printf("I am staying still\n")
		}
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Status() {
	direction := "staying still"
	if r.direction != NO_DIRECTION {
		direction = "walking " + directionToString(r.direction)
	}

	state := "off"
	if r.turnedOn {
		state = "on"
	}
	fmt.Printf("turned %s, %s\n", state, direction)
}

func directionToString(direction int64) string {
	directions := make(map[int64]string)
	directions[UP] = "up"
	directions[DOWN] = "down"
	directions[LEFT] = "left"
	directions[RIGHT] = "right"

	str, ok := directions[direction]
	if !ok {
		str = "nowhere"
	}
	return str
}

func New() *Robot {
	return &Robot{NO_DIRECTION, false}
}
