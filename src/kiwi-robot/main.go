package main

import (
	"kiwi-robot/robot"
	"flag"
	"github.com/sirupsen/logrus"
	"fmt"
	"os"
	"strings"
)

func main() {
	verbose := flag.Bool(`v`, false, `verbose output`)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [OPTIONS] COMMANDS...\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(),"Options:\n");;
		flag.PrintDefaults();;
		fmt.Fprintf(flag.CommandLine.Output(), "\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Command must be in format (PLACE X,Y,F | RIGHT | LEFT | MOVE | REPORT)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t PLACE X,Y,F - place robot to position X, Y and facing to the direction F\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t\t X,Y \t- positive or zero integer\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t\t F \t- must be one of: NORTH, EAST, SOUTH, WEST\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t LEFT \t- turn robot left\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t RIGHT \t- turn robot right\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t MOVE \t- move robot to the one position forward\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\t REPORT\t- print actual position and facing of robot in format X,Y,F\n")
	}
	flag.Parse()


	// set verbose level of log
	if verbose != nil && *verbose == true {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// build command string and clean spaces
	command := strings.Join(flag.Args(), ` `)
	command = strings.TrimSpace(command)
	if len(command) == 0 {
		logrus.Debug(`command is empty`)
	}

	// start robot
	robot.Start(command)

}