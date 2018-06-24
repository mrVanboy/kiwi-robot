package robot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	r := Robot{
		position: [2]uint8{2, 2},
		vector:   [2]int8{1, 0},
		placed:   true,
	}

	var err error

	err = move(&r).Execute()
	assert.NoError(t, err)
	// must move to the next cell
	assert.EqualValues(t, [2]uint8{3, 2}, r.position)
	assert.EqualValues(t, [2]int8{1, 0}, r.vector)

}

func testEdgeMovement(t *testing.T, robot *Robot) {
	err := move(robot).Execute()
	assert.Error(t, err)
}

func TestMoveOnEdges(t *testing.T) {
	positions := [4][2]uint8{
		{0, 0},
		{0, MAX_SIZE - 1},
		{MAX_SIZE - 1, MAX_SIZE - 1},
		{MAX_SIZE - 1, 0},
	}

	for _, pos := range positions {
		for i := range pos {
			vector := [2]int8{0, 0}

			switch pos[i] {
			case 0:
				vector[i] = -1
			case MAX_SIZE - 1:
				vector[i] = 1
			}
			r := Robot{
				position: pos,
				vector:   vector,
				placed:   true,
			}
			testEdgeMovement(t, &r)
		}
	}

}

func TestMoveOnNotPlaced(t *testing.T) {
	r := new(Robot)

	err := move(r).Execute()
	assert.Error(t, err)
}
