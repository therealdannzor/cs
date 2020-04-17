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

	// CASE 1: remove first element (1)
	s := list.New()
	exp1 := []int{2, 3, 4, 5}
	removed1 := 1

	del := &list.Element{}
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 1 {
			del = e
		}
	}
	e1 := s.Remove(del)
	res1 := listdata(s)
	assert.Equal(t, removed1, e1.Data)
	assert.Equal(t, exp1, res1)

	// CASE 2: remove mid element (3)
	s = list.New()
	exp2 := []int{1, 2, 4, 5}
	removed2 := 3
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 3 {
			del = e
		}
	}
	e2 := s.Remove(del)
	res2 := listdata(s)
	assert.Equal(t, removed2, e2.Data)
	assert.Equal(t, exp2, res2)

	// CASE 3: remove last element (5)
	s = list.New()
	exp3 := []int{1, 2, 3, 4}
	removed3 := 5
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 5 {
			del = e
		}
	}
	e3 := s.Remove(del)
	res3 := listdata(s)
	assert.Equal(t, removed3, e3.Data)
	assert.Equal(t, exp3, res3)

	// CASE 4: remove an element not in the list
	el := &list.Element{Data: 9}
	e4 := s.Remove(el)
	assert.Nil(t, e4)

	// CASE 5: remove a nil element
	e5 := s.Remove(nil)
	assert.Nil(t, e5)

	// CASE 6: remove an element from an empty list
	s = list.New()
	e6 := s.Remove(el)
	assert.Nil(t, e6)
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

	// find an element that does not exist
	exp4 := -1
	res4 := s.Find(&list.Element{Data: 9})
	assert.Equal(t, exp4, res4)

	// reset list and look for (any) element
	s = list.New()
	res5 := s.Find(firstCase)
	assert.Equal(t, exp4, res5)
}

func TestRemoveBefore(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	// new list: CASE 1 (remove the element before 2)
	s := list.New()
	item := &list.Element{}

	// expect 1 removed
	exp1 := []int{2, 3, 4, 5}
	removed1 := 1
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 2 {
			item = e
		}
	}
	e1 := s.RemoveBefore(item)
	res1 := listdata(s)
	assert.Equal(t, removed1, e1.Data)
	assert.Equal(t, exp1, res1)

	// reset list: CASE 2 (remove the element before 3)
	s = list.New()

	// expect 2 removed
	exp2 := []int{1, 3, 4, 5}
	removed2 := 2
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 3 {
			item = e
		}
	}
	e2 := s.RemoveBefore(item)
	res2 := listdata(s)
	assert.Equal(t, removed2, e2.Data)
	assert.Equal(t, exp2, res2)

	// reset list: CASE 3 (remove the element before 5)
	s = list.New()

	// expect 4 removed
	exp3 := []int{1, 2, 3, 5}
	removed3 := 4
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
		if num == 5 {
			item = e
		}
	}
	e3 := s.RemoveBefore(item)
	res3 := listdata(s)
	assert.Equal(t, removed3, e3.Data)
	assert.Equal(t, exp3, res3)

	// reset list: CASE 4 (non-nil param element which does not exist in the list)
	s = list.New()
	for _, num := range data {
		e := &list.Element{Data: num}
		s.Append(e)
	}

	exp4 := data
	el := &list.Element{Data: 9}
	e4 := s.RemoveBefore(el)
	res4 := listdata(s)
	assert.Nil(t, e4)
	assert.Equal(t, exp4, res4)

	// reset list: CASE 5 (empty list)
	s = list.New()

	e5 := s.RemoveBefore(el)
	res5 := listdata(s)
	assert.Nil(t, e5)
	assert.Nil(t, res5)

	// reset list: CASE 6 (nil param element)
	e6 := s.RemoveBefore(nil)
	res6 := listdata(s)
	assert.Nil(t, e6)
	assert.Nil(t, res6)
}

func listdata(s *list.SingleLink) []int {
	var a []int
	for e := s.First(); e != nil; e = e.Next() {
		a = append(a, e.Data)
	}

	return a
}
