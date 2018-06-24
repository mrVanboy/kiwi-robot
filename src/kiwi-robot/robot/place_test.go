package robot

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"kiwi-robot/commands"
)

func TestPlaceOutOfField(t *testing.T) {

	positions := [][2]uint8{
		{0, MAX_SIZE},
		{MAX_SIZE , MAX_SIZE },
		{MAX_SIZE, 0},
	}

	r := new(Robot)

	for _, pos := range positions {
		p := place(r)
		p.SetX(pos[0])
		p.SetY(pos[1])
		p.SetDir(`NORTH`)
		assert.EqualValues(t, ErrPositionOutOfField, p.Execute())
	}

}

func TestPlaceBadDirection(t *testing.T){
	r := new(Robot)

	p := place(r)
	p.SetX(2)
	p.SetY(3)
	p.SetDir(`Adabra-kedabra`)

	err := p.Execute()
	assert.Error(t, err)
	assert.EqualValues(t, ErrBadDirection, err)
}

func TestPlaceOnNotPlaced(t *testing.T) {
	r := new(Robot)

	p := place(r)
	p.SetX(2)
	p.SetY(3)
	p.SetDir(commands.DIR_NORTH)

	err := p.Execute()
	assert.NoError(t, err)

	assert.EqualValues(t, r.vector, [2]int8{0,1})
	assert.EqualValues(t, r.position, [2]uint8{2,3})
	assert.True(t, r.placed)
}
