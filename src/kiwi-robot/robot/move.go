package robot

import (
	"errors"
)

var ErrTargetPositionIsOutOfField = errors.New(`target position is out of field`)

// cmdMove implements ICommand interface and move robot to the one position
// forward in the direction, where it is faced.
type cmdMove struct {
	robot *Robot
}

func (c *cmdMove) Execute() error {
	// skip execution if robot is not placed
	if  c.robot.placed == false {
		return ErrNotPlaced
	}

	newPosition := c.robot.position

	for i, value := range c.robot.vector {
		// check if new position is not out of field
		if temp := int8(newPosition[i]) + value; temp < MAX_SIZE && temp >= 0 {
			newPosition[i] = uint8(temp)
		} else {
			return ErrTargetPositionIsOutOfField
		}
	}

	c.robot.position = newPosition

	return nil
}

// move function create new pointer to the cmdMove struct
func move(robot *Robot) *cmdMove {
	return &cmdMove{robot:robot}
}

