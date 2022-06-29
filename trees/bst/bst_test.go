package bst_test

import (
	"testing"

	bstree "github.com/diegolopes98/go-data-structs/trees/bst"
	"github.com/stretchr/testify/assert"
)

func TestEmptyBSTInsert(t *testing.T) {
	bst := bstree.New[int]()
	want := 10
	wantfreq := uint(1)
	bst.Insert(want)
	assert.Equal(t, want, bst.Root().Value)
	assert.Equal(t, wantfreq, bst.Root().Frequency)
}

func TestBSTInsertToRight(t *testing.T) {
	bst := bstree.New[int]()
	want := 10
	wantfreq := uint(1)
	bst.Insert(1).Insert(10)
	assert.Equal(t, want, bst.Root().Right.Value)
	assert.Equal(t, wantfreq, bst.Root().Right.Frequency)
}

func TestBSTInsertToRightDeeper(t *testing.T) {
	bst := bstree.New[int]()
	want := 100
	wantfreq := uint(1)
	bst.Insert(1).Insert(10).Insert(100)
	assert.Equal(t, want, bst.Root().Right.Right.Value)
	assert.Equal(t, wantfreq, bst.Root().Right.Right.Frequency)
}

func TestBSTInsertToLeft(t *testing.T) {
	bst := bstree.New[int]()
	want := 10
	wantfreq := uint(1)
	bst.Insert(100).Insert(10)
	assert.Equal(t, want, bst.Root().Left.Value)
	assert.Equal(t, wantfreq, bst.Root().Left.Frequency)
}

func TestBSTInsertToLeftDeeper(t *testing.T) {
	bst := bstree.New[int]()
	want := 1
	wantfreq := uint(1)
	bst.Insert(100).Insert(10).Insert(1)
	assert.Equal(t, want, bst.Root().Left.Left.Value)
	assert.Equal(t, wantfreq, bst.Root().Left.Left.Frequency)
}

func TestBSTInsertToFreq(t *testing.T) {
	bst := bstree.New[int]()
	want := 10
	wantfreq := uint(2)
	bst.Insert(10).Insert(10)
	assert.Equal(t, want, bst.Root().Value)
	assert.Equal(t, wantfreq, bst.Root().Frequency)
}
