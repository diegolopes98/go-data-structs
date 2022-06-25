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
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Pop()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errv)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestSllPopOneElem(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := 1
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Pop()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestSllPop(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2)
	wanthead := 1
	wanttail := 2
	wantlen := uint(1)
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Pop()
	len := sll.Length()
	assert.Equal(t, wanttail, value)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantlen, len)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestSllShiftEmpty(t *testing.T) {
	sll := New[int]()
	want := 0
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Shift()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errv)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestSllShiftOneElem(t *testing.T) {
	sll := New[int]()
	sll.Push(1)
	want := 1
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Shift()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestSllShift(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2)
	wanthead := 1
	wanttail := 2
	wantlen := uint(1)
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	value, errv := sll.Shift()
	len := sll.Length()
	assert.Equal(t, wanthead, value)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, len)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestSllUnshiftEmpty(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1)
	want := 1
	head, errhead := sll.Head()
	tail, errtail := sll.Tail()
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestSllUnshift(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wanttail := 1
	wantlen := uint(3)
	head, errhead := sll.Head()
	tail, errtail := sll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, sll.Length())
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestSllGetOutOfBounds(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1).Unshift(2).Unshift(3)
	want := 0
	value, err := sll.Get(100)
	assert.Equal(t, want, value)
	assert.NotNil(t, err)
}

func TestSllGet(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1).Unshift(2).Unshift(3)
	want := 1
	value, err := sll.Get(2)
	assert.Equal(t, want, value)
	assert.Nil(t, err)
}

func TestSllSetOutOfBounds(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wantmid := 2
	wanttail := 1
	sll.Set(100, 100)
	head, errh := sll.Head()
	mid, errm := sll.Get(1)
	tail, errt := sll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantmid, mid)
	assert.Equal(t, wanttail, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errm)
	assert.Nil(t, errt)
}

func TestSllSet(t *testing.T) {
	sll := New[int]()
	sll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wantmid := 100
	wanttail := 1
	sll.Set(1, 100)
	head, errh := sll.Head()
	mid, errm := sll.Get(1)
	tail, errt := sll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantmid, mid)
	assert.Equal(t, wanttail, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errm)
	assert.Nil(t, errt)
}

func TestSllInsertAtBeginning(t *testing.T) {
	sll := New[int]()
	wantoldlen := uint(0)
	wantnewlen := uint(1)
	want := 10
	oldlen := sll.Length()
	sll.Insert(0, 10)
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	newlen := sll.Length()
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestSllInsertOutOfBounds(t *testing.T) {
	sll := New[int]()
	wantoldlen := uint(0)
	wantnewlen := uint(0)
	want := 0
	oldlen := sll.Length()
	sll.Insert(1000, 10)
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	newlen := sll.Length()
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestSllInsert(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(3).Push(4)
	wantoldlen := uint(3)
	wantnewlen := uint(4)
	want := 2
	wantnext := 3
	oldlen := sll.Length()
	sll.Insert(1, 2)
	newlen := sll.Length()
	value, errv := sll.Get(1)
	nextvalue, errnv := sll.Get(2)
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, value)
	assert.Equal(t, wantnext, nextvalue)
	assert.Nil(t, errv)
	assert.Nil(t, errnv)
}

func TestSllRemoveEmpty(t *testing.T) {
	sll := New[int]()
	want := 0
	value, errv := sll.Remove(1)
	assert.Equal(t, want, value)
	assert.NotNil(t, errv)
}

func TestSllRemoveFromTheBeginning(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2)
	want := 1
	value, errv := sll.Remove(0)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestSllRemoveFromTheEnd(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2)
	want := 2
	value, errv := sll.Remove(1)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestSllRemoveFromTheMid(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2).Push(3)
	want := 2
	value, errv := sll.Remove(1)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestSllReverse(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2).Push(3)
	wanthead := 3
	wanttail := 1
	wantmid := 2
	sll.Reverse()
	head, errh := sll.Head()
	tail, errt := sll.Tail()
	mid, errm := sll.Get(1)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantmid, mid)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
	assert.Nil(t, errm)
}

func TestSllForEach(t *testing.T) {
	sll := New[int]()
	sll.Push(1).Push(2).Push(3)
	want := []int{1, 2, 3}
	sll.ForEach(func(number *int) {
		assert.Contains(t, want, *number)
	})
}
