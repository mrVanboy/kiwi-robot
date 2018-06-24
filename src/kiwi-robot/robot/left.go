package robot

// cmdLeft struct implements ICommand interface and turn robot left at execution
type cmdLeft struct {
	robot *Robot
}

func (c *cmdLeft) Execute() error {
	// skip execution if robot is not placed
	if  c.robot.placed == false {
		return ErrNotPlaced
	}

	tmp := c.robot.vector[0]
	c.robot.vector[0] = -1 * c.robot.vector[1]
	c.robot.vector[1] = tmp

	return nil
}

// left function create new pointer to the cmdLeft struct
func left(robot *Robot) *cmdLeft {
	return &cmdLeft{robot:robot}
}

