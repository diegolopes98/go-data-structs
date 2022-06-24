package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSllSize(t *testing.T) {
	sll := New[int]()
	wantlen := uint(0)
	length := sll.Length()
	assert.Equal(t, wantlen, length)
}

func TestNewSllHead(t *testing.T) {
	sll := New[int]()
	want := 0
	head, err := sll.Head()
	assert.Equal(t, head, want)
	assert.NotNil(t, err)
}

func TestNewSllTail(t *testing.T) {
	sll := New[int]()
	want := 0
	tail, err := sll.Tail()
	assert.Equal(t, tail, want)
	assert.NotNil(t, err)
}

func TestNewSllPush(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := 1
	head, errhead := sll.Head()
	tail, errtail := sll.Tail()
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}
