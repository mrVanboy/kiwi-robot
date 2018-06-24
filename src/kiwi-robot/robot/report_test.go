package robot

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	r := Robot{
		position: [2]uint8{2,3},
		vector:   [2]int8{1,0},
		placed:   true,
	}

	var buf bytes.Buffer
	oldOutput := output
	output = &buf
	assert.NoError(t, report(&r).Execute())
	output = oldOutput

	result := buf.String()
	assert.EqualValues(t, "2,3,EAST\n", result)
}

func TestCmdReportNotPlaced(t *testing.T) {
	r := new(Robot)
	assert.EqualError(t, report(r).Execute(), ErrNotPlaced.Error())
}
