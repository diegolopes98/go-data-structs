package stack_test

import (
	"testing"

	"github.com/diegolopes98/go-data-structs/stack"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStackSize(t *testing.T) {
	stck := stack.New[int]()
	want := uint(0)
	assert.Equal(t, want, stck.Size())
}

func TestEmptyStackPop(t *testing.T) {
	stck := stack.New[int]()
	value, err := stck.Pop()
	want := 0
	assert.Equal(t, want, value)
	assert.NotNil(t, err)
}

func TestEmptyStackPushAndPop(t *testing.T) {
	stck := stack.New[int]()
	stck.Push(10)
	actualsize := stck.Size()
	value, err := stck.Pop()
	wantsize := uint(1)
	wantvalue := 10
	assert.Equal(t, wantsize, actualsize)
	assert.Equal(t, wantvalue, value)
	assert.Nil(t, err)
}

func TestStackPushAndPop(t *testing.T) {
	stck := stack.New[int]()
	stck.Push(10).Push(20)
	actualsize := stck.Size()
	value, err := stck.Pop()
	wantsize := uint(2)
	wantvalue := 20
	assert.Equal(t, wantsize, actualsize)
	assert.Equal(t, wantvalue, value)
	assert.Nil(t, err)
}
