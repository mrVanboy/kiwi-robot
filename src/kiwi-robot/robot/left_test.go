package robot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	r := Robot{
		position: [2]uint8{2, 2},
		vector:   [2]int8{1, 0},
		placed:   true,
	}

	var err error

	err = left(&r).Execute()
	assert.NoError(t, err)
	// must be north
	assert.EqualValues(t, [2]int8{0, 1}, r.vector)
	assert.EqualValues(t, [2]uint8{2, 2}, r.position)

	err = left(&r).Execute()
	assert.NoError(t, err)
	// must be west
	assert.EqualValues(t, [2]int8{-1, 0}, r.vector)
	assert.EqualValues(t, [2]uint8{2, 2}, r.position)

	err = left(&r).Execute()
	assert.NoError(t, err)
	// must be south
	assert.EqualValues(t, [2]int8{0, -1}, r.vector)
	assert.EqualValues(t, [2]uint8{2, 2}, r.position)

	err = left(&r).Execute()
	assert.NoError(t, err)
	// must be east
	assert.EqualValues(t, [2]int8{1, 0}, r.vector)
	assert.EqualValues(t, [2]uint8{2, 2}, r.position)

}

func TestLeftOnNotPlaced(t *testing.T) {
	r := new(Robot)

	err := left(r).Execute()
	assert.Error(t, err)
}
