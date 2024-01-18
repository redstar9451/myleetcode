package p0641

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP0641(t *testing.T) {
	obj := Constructor(3)
	assert.Equal(t, true, obj.InsertFront(8))
	assert.Equal(t, true, obj.InsertLast(8))
	assert.Equal(t, true, obj.InsertLast(2))
	assert.Equal(t, 8, obj.GetFront())
	assert.Equal(t, true, obj.DeleteLast())
	assert.Equal(t, 8, obj.GetRear())
	assert.Equal(t, true, obj.InsertFront(9))
	assert.Equal(t, true, obj.DeleteFront())
	assert.Equal(t, 8, obj.GetRear())
	assert.Equal(t, true, obj.InsertLast(2))
	assert.Equal(t, true, obj.IsFull())
}
