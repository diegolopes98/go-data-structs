package queue_test

import (
	"testing"

	"github.com/diegolopes98/go-data-structs/queue"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStackSize(t *testing.T) {
	stck := queue.New[int]()
	want := uint(0)
	assert.Equal(t, want, stck.Size())
}

func TestEmptyStackDequeue(t *testing.T) {
	stck := queue.New[int]()
	value, err := stck.Dequeue()
	want := 0
	assert.Equal(t, want, value)
	assert.NotNil(t, err)
}

func TestEmptyStackEnqueueAndDequeue(t *testing.T) {
	stck := queue.New[int]()
	stck.Enqueue(10)
	actualsize := stck.Size()
	value, err := stck.Dequeue()
	wantsize := uint(1)
	wantvalue := 10
	assert.Equal(t, wantsize, actualsize)
	assert.Equal(t, wantvalue, value)
	assert.Nil(t, err)
}

func TestStackEnqueueAndDequeue(t *testing.T) {
	stck := queue.New[int]()
	stck.Enqueue(10).Enqueue(20)
	actualsize := stck.Size()
	value, err := stck.Dequeue()
	wantsize := uint(2)
	wantvalue := 10
	assert.Equal(t, wantsize, actualsize)
	assert.Equal(t, wantvalue, value)
	assert.Nil(t, err)
}
