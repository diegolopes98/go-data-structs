package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSllLength(t *testing.T) {
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

func TestSllLength(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := uint(1)
	assert.Equal(t, want, sll.Length())
}

func TestSllHead(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := 1
	head, err := sll.Head()
	assert.Equal(t, head, want)
	assert.Nil(t, err)
}

func TestSllTail(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := 1
	tail, err := sll.Tail()
	assert.Equal(t, tail, want)
	assert.Nil(t, err)
}

func TestSllPush(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2).Push(3)
	wanthead := 1
	wanttail := 3
	wantlen := uint(3)
	head, errhead := sll.Head()
	tail, errtail := sll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, sll.Length())
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestSllPopEmpty(t *testing.T) {
	sll := New[int]()
	want := 0
	value, errv := sll.Pop()
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errv)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}
