package mapslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	ID    int
	Name  string
	Money float64
	Male  bool
}

type Group struct {
	People []*Person
}

func TestMapSliceToString(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}
	s, err := ToStrings(g.People, "Name")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], "George")
	assert.Equal(s[1], "Jeff")
	assert.Equal(s[2], "Ted")
	assert.Equal(s[3], "Luda")

	s, err = ToStrings(g.People, "ID")
	assert.Equal(ErrNotString, err)
}

func TestMapSliceToInt(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := ToInts(g.People, "ID")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], 0)
	assert.Equal(s[1], 1)
	assert.Equal(s[2], 2)
	assert.Equal(s[3], 3)

	s, err = ToInts(g.People, "Money")
	assert.Equal(ErrNotInt, err)
}

func TestMapSliceToFloat(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := ToFloats(g.People, "Money")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], 42.42)
	assert.Equal(s[1], 0.0)
	assert.Equal(s[2], 50.0)
	assert.Equal(s[3], 100.0)

	s, err = ToFloats(g.People, "ID")
	assert.Equal(ErrNotFloat, err)
}

func TestMapSliceToBool(t *testing.T) {
	assert := assert.New(t)

	g := &Group{
		People: []*Person{
			{0, "George", 42.42, true},
			{1, "Jeff", 0, true},
			{2, "Ted", 50, true},
			{3, "Luda", 100, false},
		},
	}

	s, err := ToBools(g.People, "Male")
	assert.Nil(err)
	assert.NotEmpty(s)
	assert.Equal(s[0], true)
	assert.Equal(s[1], true)
	assert.Equal(s[2], true)
	assert.Equal(s[3], false)

	s, err = ToBools(g.People, "Name")
	assert.Equal(ErrNotBool, err)
}
