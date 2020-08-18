package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	position := Position{}
	var movement NewMovement = &MovementAdapter{OldMovement{}}

	assert.Equal(t, 0, position.x)
	assert.Equal(t, 0, position.y)

	movement.leftDown(&position)
	assert.Equal(t, -1, position.x)
	assert.Equal(t, -1, position.y)

	movement.leftUp(&position)
	assert.Equal(t, -2, position.x)
	assert.Equal(t, 0, position.y)

	movement.rightDown(&position)
	assert.Equal(t, -1, position.x)
	assert.Equal(t, -1, position.y)

	movement.rightUp(&position)
	assert.Equal(t, 0, position.x)
	assert.Equal(t, 0, position.y)
}
