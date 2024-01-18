package p0232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP0232(t *testing.T) {
	q := Constructor()
	assert.Equal(t, true, q.Empty())
	q.Push(1)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, false, q.Empty())
	q.Push(2)
	q.Push(3)
	q.Pop()
	assert.Equal(t, 2, q.Peek())
	q.Push(4)
	assert.Equal(t, false, q.Empty())
}
