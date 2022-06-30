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

func TestBSTEmptyFind(t *testing.T) {
	bst := bstree.New[int]()
	node := bst.Find(10)
	assert.Nil(t, node)
}

func TestBSTFindRoot(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(10)
	want := 10
	node := bst.Find(10)
	assert.NotNil(t, node)
	assert.Equal(t, want, node.Value)
}

func TestBSTFindRight(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(1).Insert(10)
	want := 10
	node := bst.Find(10)
	assert.NotNil(t, node)
	assert.Equal(t, want, node.Value)
}

func TestBSTFindRightDeeper(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(1).Insert(10).Insert(100)
	want := 100
	node := bst.Find(100)
	assert.NotNil(t, node)
	assert.Equal(t, want, node.Value)
}

func TestBSTFindLeft(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(10).Insert(1)
	want := 1
	node := bst.Find(1)
	assert.NotNil(t, node)
	assert.Equal(t, want, node.Value)
}

func TestBSTFindLeftDeeper(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(100).Insert(10).Insert(1)
	want := 1
	node := bst.Find(1)
	assert.NotNil(t, node)
	assert.Equal(t, want, node.Value)
}

func TestBSTFindNone(t *testing.T) {
	bst := bstree.New[int]()
	bst.Insert(100).Insert(10).Insert(1)
	node := bst.Find(1000)
	assert.Nil(t, node)
}
