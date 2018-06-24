package robot

import (
	"kiwi-robot/commands"
	log "github.com/sirupsen/logrus"
	"errors"
)

const MAX_SIZE = 5
var ErrNotPlaced = errors.New(`robot is not placed, skipping command`)
type Robot struct {
	position [2]uint8
	vector [2]int8
	placed bool
}

// Start function create new robot and consume sequence if commands in sCommand
// string. Ex.:
//		robot.Start(`PLACE 1,0,EAST MOVE LEFT REPORT`)
func Start(sCommands string){
	r := new(Robot)

	// create new config for commands
	cfg := commands.Config{
		Move:   move(r),
		Left:   left(r),
		Right:  right(r),
		Report: report(r),
		Place: place(r),
	}

	// get parsed commands array
	cmds, err := commands.ParseCommandString(sCommands, cfg)
	if err != nil {
		log.Fatal(`command error: `, err)
	}

	// execute commands
	for _, v := range cmds {
		if v == nil {
			continue
		}

		if err := v.Execute(); err != nil {
			log.Debug(err)
		}
	}
}
