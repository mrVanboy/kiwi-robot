package commands

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type right struct {
}

func (*right) Execute() error {
	return nil
}

type left struct {
}

func (*left) Execute() error {
	return nil
}

type move struct {
}

func (*move) Execute() error {
	return nil
}

type report struct {
}

func (*report) Execute() error {
	return nil
}

type place struct {}

func (*place) Clone() IPlacer {
	return new(place)
}

func (*place) SetX(x uint8) {
}

func (*place) SetY(y uint8) {
}

func (*place) SetDir(dir string) {
}

func (*place) Execute() error {
	return nil
}

type CommandTestSuite struct {
	suite.Suite
	cfg Config
}

func (suite *CommandTestSuite) SetupTest() {
	suite.cfg = Config{
		Right: new(right),
		Left: new(left),
		Move: new(move),
		Report: new(report),
		Place: new(place),
	}
}

func (suite *CommandTestSuite) TestMove() {
	fixtures := map[string][]ICommand{
		`PLACE 0,0,NORTH MOVE REPORT`: {suite.cfg.Place, suite.cfg.Move, suite.cfg.Report},
		`PLACE 0,0,NORTH LEFT REPORT`: {suite.cfg.Place, suite.cfg.Left, suite.cfg.Report},
		`PLACE 1,2,EAST MOVE MOVE LEFT MOVE REPORT`: {suite.cfg.Place, suite.cfg.Move, suite.cfg.Move, suite.cfg.Left, suite.cfg.Move, suite.cfg.Report},
	}

	for str, cmds := range fixtures {
		actualCmds, err := ParseCommandString(str, suite.cfg)
		suite.NoError(err)
		suite.Equal(cmds, actualCmds)
	}
}

func TestCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}