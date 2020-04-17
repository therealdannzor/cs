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
	expected := []int{1, 2, 3, 4, 5}

	s := list.New()
	data := []int{2, 3, 4, 5}
	for _, v := range data {
		s.Append(&list.Element{Data: v})
	}
	s.Prepend(&list.Element{Data: 1})

	actual := listdata(s)

	assert.Equal(t, expected, actual)
}

func TestRemove(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := list.New()
	del := &list.Element{}

	// remove first element
	// expect 1 removed
	exp1 := []int{2, 3, 4, 5}
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 1 {
			del = e
		}
	}
	s.Remove(del)
	res1 := listdata(s)
	assert.Equal(t, exp1, res1)

	// reset list
	s = list.New()

	// remove mid element
	// expect 3 removed
	exp2 := []int{1, 2, 4, 5}
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 3 {
			del = e
		}
	}
	s.Remove(del)
	res2 := listdata(s)
	assert.Equal(t, exp2, res2)

	// reset list
	s = list.New()

	// remove last element
	// expect 5 removed
	exp3 := []int{1, 2, 3, 4}
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 5 {
			del = e
		}
	}
	s.Remove(del)
	res3 := listdata(s)
	assert.Equal(t, exp3, res3)
}

func TestFind(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := list.New()
	firstCase := &list.Element{}
	secondCase := &list.Element{}
	thirdCase := &list.Element{}

	// find first element (1)
	// expect index 0
	exp1 := 0
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 1 {
			firstCase = e
		}
		if num == 3 {
			secondCase = e
		}
		if num == 5 {
			thirdCase = e
		}
	}
	res1 := s.Find(firstCase)
	assert.Equal(t, exp1, res1)

	// find mid element (3)
	// expect index 2
	exp2 := 2
	res2 := s.Find(secondCase)
	assert.Equal(t, exp2, res2)

	// find last element (5)
	// expect index 4
	exp3 := 4
	res3 := s.Find(thirdCase)
	assert.Equal(t, exp3, res3)
}

func listdata(s *list.SingleLink) []int {
	var a []int
	for e := s.First(); e != nil; e = e.Next() {
		a = append(a, e.Data)
	}

	return a
}
