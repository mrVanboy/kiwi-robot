package commands

import (
	"errors"
	"regexp"
	"strings"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	ErrBadString = errors.New(`can't parse command string`)
)

// Config struct holding callbacks for invoking commands
type Config struct {
	Move, Left, Right, Report ICommand
	Place                     IPlacer
}

// ParseCommandString function parse sCommand string with sequence of commands.
// cmdsCallback must contain callbacks for returning right sequence of parsed
// commands
func ParseCommandString(sCommands string, cmdsCallback Config) ([]ICommand, error) {
	commands := make([]ICommand, 0)
	r := regexp.MustCompile(mainRegexpTemplate)

	// prepare variable for checking if parsing was successful
	replaced := true
	for replaced {
		replaced = false
		sCommands = r.ReplaceAllStringFunc(sCommands, func(match string) string {
			replaced = true
			commands = append(commands, parseSingleCommand(strings.TrimSpace(match), cmdsCallback))
			return ``
		})
	}

	// check if command string contain the unparsed rest
	if len(sCommands) > 0 {
		return nil, ErrBadString
	}

	return commands, nil
}

// parseSingleCommand function parse one command like REPORT, MOVE etc. and return
// appropriate callback
func parseSingleCommand(sCommand string, cmdsCallback Config) ICommand {
	switch sCommand{
	case CMD_REPORT:
		return cmdsCallback.Report
	case CMD_RIGHT:
		return cmdsCallback.Right
	case CMD_LEFT:
		return cmdsCallback.Left
	case CMD_MOVE:
		return cmdsCallback.Move
	}
	r := regexp.MustCompile(placeRegexp)
	matches := r.FindStringSubmatch(sCommand)
	if matches == nil {
		log.Debug(`unknown command ` + sCommand)
		return nil
	}
	c := cmdsCallback.Place
	for i, name := range r.SubexpNames() {
		switch name {
		case rNamePlaceX:
			x, err := strconv.ParseUint(matches[i], 10, 8)
			if err != nil {
				log.Debug(`can't parse x value ` + matches[i])
				return nil
			}
			c.SetX(uint8(x))
		case rNamePlaceY:
			x, err := strconv.ParseUint(matches[i], 10, 8)
			if err != nil {
				log.Debug(`can't parse y value ` + matches[i])
				return nil
			}
			c.SetY(uint8(x))
		case rNamePlaceDir:
			c.SetDir(matches[i])
		}
	}
	return c
}
