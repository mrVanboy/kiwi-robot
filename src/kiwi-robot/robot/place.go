package robot

import (
	"errors"
)

// directionNameMap holds inverted map of directionMap
var directionNameMap = func()map[string][2]int8 {
	newMap := make(map[string][2]int8, len(directionMap))
	for k, v := range directionMap {
		newMap[v] = k
	}
	return newMap
}()

var (
	ErrPositionOutOfField = errors.New("position is out of field")
	ErrBadDirection = errors.New(`can't recognize direction'`)
)

// cmdPlace struct implements IPlacer interface and place robot to the new
// position at execution
type cmdPlace struct {
	x, y uint8
	dir string
	robot *Robot
}

func (c *cmdPlace) SetX(x uint8) {
	c.x = x
}

func (c *cmdPlace) SetY(y uint8) {
	c.y = y
}

func (c *cmdPlace) SetDir(dir string) {
	c.dir = dir
}

func (c *cmdPlace) Execute() error {
	if c.x >= MAX_SIZE || c.y >= MAX_SIZE || c.x < 0 || c.y < 0 {
		return ErrPositionOutOfField
	}

	if vector, ok := directionNameMap[c.dir]; !ok{
		return ErrBadDirection
	} else {
		c.robot.vector = vector
	}

	c.robot.position = [2]uint8{c.x, c.y}
	c.robot.placed = true

	return nil
}

// place function create new pointer to the cmdPlace struct
func place(robot *Robot) *cmdPlace {
	return &cmdPlace{robot:robot}
}
