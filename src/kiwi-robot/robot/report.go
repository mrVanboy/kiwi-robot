package robot

import (
	"fmt"
	"kiwi-robot/commands"
	"os"
	"io"
)

const REPORT_TEMPLATE = "%d,%d,%s"

// directionMap variable hold map of directions strings with keys as directions vectors
var directionMap = map[[2]int8]string{
	[2]int8{0,1}: commands.DIR_NORTH,
	[2]int8{0,-1}: commands.DIR_SOUTH,
	[2]int8{1,0}: commands.DIR_EAST,
	[2]int8{-1,0}: commands.DIR_WEST,
	[2]int8{0,0}: `UNKNOWN`,
}

var output io.Writer = os.Stdout

// cmdReport struct implements ICommand interface and print current position and
// facing direction of the robot
type cmdReport struct {
	robot *Robot
}

func (c *cmdReport) Execute() error{
	if  c.robot.placed == false {
		return ErrNotPlaced
	}
	fmt.Fprintf(output, REPORT_TEMPLATE + "\n", c.robot.position[0], c.robot.position[1], directionMap[c.robot.vector])
	return nil
}

// report function create pointer to the cmdReport struct
func report(robot *Robot) *cmdReport {
	return &cmdReport{robot:robot}
}
