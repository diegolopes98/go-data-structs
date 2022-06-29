package dll_test

import (
	"testing"

	doubly "github.com/diegolopes98/go-data-structs/lists/dll"
	"github.com/stretchr/testify/assert"
)

func TestNewDllLength(t *testing.T) {
	dll := doubly.New[int]()
	wantlen := uint(0)
	length := dll.Length()
	assert.Equal(t, wantlen, length)
}

func TestNewDllHead(t *testing.T) {
	dll := doubly.New[int]()
	want := 0
	head, err := dll.Head()
	assert.Equal(t, head, want)
	assert.NotNil(t, err)
}

func TestNewDllTail(t *testing.T) {
	dll := doubly.New[int]()
	want := 0
	tail, err := dll.Tail()
	assert.Equal(t, tail, want)
	assert.NotNil(t, err)
}

func TestNewDllPush(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := 1
	head, errhead := dll.Head()
	tail, errtail := dll.Tail()
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestDllLength(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := uint(1)
	assert.Equal(t, want, dll.Length())
}

func TestDllHead(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := 1
	head, err := dll.Head()
	assert.Equal(t, head, want)
	assert.Nil(t, err)
}

func TestDllTail(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := 1
	tail, err := dll.Tail()
	assert.Equal(t, tail, want)
	assert.Nil(t, err)
}

func TestDllPush(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3)
	wanthead := 1
	wanttail := 3
	wantlen := uint(3)
	head, errhead := dll.Head()
	tail, errtail := dll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, dll.Length())
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestDllPopEmpty(t *testing.T) {
	dll := doubly.New[int]()
	want := 0
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Pop()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errv)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestDllPopOneElem(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := 1
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Pop()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestDllPop(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2)
	wanthead := 1
	wanttail := 2
	wantlen := uint(1)
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Pop()
	len := dll.Length()
	assert.Equal(t, wanttail, value)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantlen, len)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestDllShiftEmpty(t *testing.T) {
	dll := doubly.New[int]()
	want := 0
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Shift()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errv)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestDllShiftOneElem(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1)
	want := 1
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Shift()
	assert.Equal(t, want, value)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestDllShift(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2)
	wanthead := 1
	wanttail := 2
	wantlen := uint(1)
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	value, errv := dll.Shift()
	len := dll.Length()
	assert.Equal(t, wanthead, value)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, len)
	assert.Nil(t, errv)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestDllUnshiftEmpty(t *testing.T) {
	dll := doubly.New[int]()
	dll.Unshift(1)
	want := 1
	head, errhead := dll.Head()
	tail, errtail := dll.Tail()
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestDllUnshift(t *testing.T) {
	dll := doubly.New[int]()
	dll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wanttail := 1
	wantlen := uint(3)
	head, errhead := dll.Head()
	tail, errtail := dll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantlen, dll.Length())
	assert.Nil(t, errhead)
	assert.Nil(t, errtail)
}

func TestDllGetOutOfBounds(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3)
	want := 0
	value, err := dll.Get(100)
	assert.Equal(t, want, value)
	assert.NotNil(t, err)
}

func TestDllGetNearBeginning(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3).Push(4)
	want := 2
	value, err := dll.Get(1)
	assert.Equal(t, want, value)
	assert.Nil(t, err)
}

func TestDllGetNearEnd(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3).Push(4)
	want := 3
	value, err := dll.Get(2)
	assert.Equal(t, want, value)
	assert.Nil(t, err)
}

func TestDllSetOutOfBounds(t *testing.T) {
	dll := doubly.New[int]()
	dll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wantmid := 2
	wanttail := 1
	dll.Set(100, 100)
	head, errh := dll.Head()
	mid, errm := dll.Get(1)
	tail, errt := dll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantmid, mid)
	assert.Equal(t, wanttail, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errm)
	assert.Nil(t, errt)
}

func TestDllSet(t *testing.T) {
	dll := doubly.New[int]()
	dll.Unshift(1).Unshift(2).Unshift(3)
	wanthead := 3
	wantmid := 100
	wanttail := 1
	dll.Set(1, 100)
	head, errh := dll.Head()
	mid, errm := dll.Get(1)
	tail, errt := dll.Tail()
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wantmid, mid)
	assert.Equal(t, wanttail, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errm)
	assert.Nil(t, errt)
}

func TestDllInsertAtBeginning(t *testing.T) {
	dll := doubly.New[int]()
	wantoldlen := uint(0)
	wantnewlen := uint(1)
	want := 10
	oldlen := dll.Length()
	dll.Insert(0, 10)
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	newlen := dll.Length()
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
}

func TestDllInsertOutOfBounds(t *testing.T) {
	dll := doubly.New[int]()
	wantoldlen := uint(0)
	wantnewlen := uint(0)
	want := 0
	oldlen := dll.Length()
	dll.Insert(1000, 10)
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	newlen := dll.Length()
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, head)
	assert.Equal(t, want, tail)
	assert.NotNil(t, errh)
	assert.NotNil(t, errt)
}

func TestDllInsert(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(3).Push(4)
	wantoldlen := uint(3)
	wantnewlen := uint(4)
	want := 2
	wantnext := 3
	oldlen := dll.Length()
	dll.Insert(1, 2)
	newlen := dll.Length()
	value, errv := dll.Get(1)
	nextvalue, errnv := dll.Get(2)
	assert.Equal(t, wantoldlen, oldlen)
	assert.Equal(t, wantnewlen, newlen)
	assert.Equal(t, want, value)
	assert.Equal(t, wantnext, nextvalue)
	assert.Nil(t, errv)
	assert.Nil(t, errnv)
}

func TestDllRemoveEmpty(t *testing.T) {
	dll := doubly.New[int]()
	want := 0
	value, errv := dll.Remove(1)
	assert.Equal(t, want, value)
	assert.NotNil(t, errv)
}

func TestDllRemoveFromTheBeginning(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2)
	want := 1
	value, errv := dll.Remove(0)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestDllRemoveFromTheEnd(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2)
	want := 2
	value, errv := dll.Remove(1)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestDllRemoveFromTheMid(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3)
	want := 2
	value, errv := dll.Remove(1)
	assert.Equal(t, want, value)
	assert.Nil(t, errv)
}

func TestDllReverse(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3)
	wanthead := 3
	wanttail := 1
	wantmid := 2
	dll.Reverse()
	head, errh := dll.Head()
	tail, errt := dll.Tail()
	mid, errm := dll.Get(1)
	assert.Equal(t, wanthead, head)
	assert.Equal(t, wanttail, tail)
	assert.Equal(t, wantmid, mid)
	assert.Nil(t, errh)
	assert.Nil(t, errt)
	assert.Nil(t, errm)
}

func TestDllForEach(t *testing.T) {
	dll := doubly.New[int]()
	dll.Push(1).Push(2).Push(3)
	want := []int{1, 2, 3}
	dll.ForEach(func(number *int) {
		assert.Contains(t, want, *number)
	})
}
