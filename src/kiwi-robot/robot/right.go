package robot

// cmdRight struct implements ICommand interface and turns robot right at
// execution
type cmdRight struct {
	robot *Robot
}

func (c *cmdRight) Execute() error {
	// skip execution if robot is not placed
	if  c.robot.placed == false {
		return ErrNotPlaced
	}

	tmp := c.robot.vector[0]
	c.robot.vector[0] = c.robot.vector[1]
	c.robot.vector[1] = -1 * tmp

	return nil
}

// right function create and return new pointer to the cmdRight struct
func right(robot *Robot) *cmdRight {
	return &cmdRight{robot:robot}
}

