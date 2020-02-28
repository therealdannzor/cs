package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/therealdannzor/cs/list"
)

func TestFirst(t *testing.T) {
	expected := &list.Element{Data: 1}

	s := list.New()
	s.Append(&list.Element{Data: 1})
	actual := s.First()

	assert.Equal(t, expected, actual)
}

func TestAppend(t *testing.T) {
	expected := []int{1, 2, 3}

	s := list.New()
	for i := 1; i < 4; i++ {
		s.Append(&list.Element{Data: i})
	}

	actual := listdata(s)

	assert.Equal(t, expected, actual)
}

func TestPrepend(t *testing.T) {
	expected := []int{1, 2, 3}

	s := list.New()
	s.Append(&list.Element{Data: 2})
	s.Append(&list.Element{Data: 3})
	s.Prepend(&list.Element{Data: 1})

	actual := listdata(s)

	assert.Equal(t, expected, actual)
}

func TestRemove(t *testing.T) {
	expected := []int{1, 2, 4}

	s := list.New()
	s.Append(&list.Element{Data: 1})
	s.Append(&list.Element{Data: 2})
	toDelete := &list.Element{Data: 3}
	s.Append(toDelete)
	s.Append(&list.Element{Data: 4})
	s.Remove(toDelete)

	actual := listdata(s)

	assert.Equal(t, expected, actual)
}

func listdata(s *list.SingleLink) []int {
	var a []int
	for e := s.First(); e != nil; e = e.Next() {
		a = append(a, e.Data)
	}

	return a
}
