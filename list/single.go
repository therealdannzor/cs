package list

// Element is each node in the linked list and has a Data field and
// a link to the next element in the list
type Element struct {
	Data int
	next *Element
}

func (e *Element) Next() *Element {
	if p := e.next; p != nil {
		return p
	}
	return nil
}

type SingleLink struct {
	head *Element
	len  int
}

func (s *SingleLink) Init() *SingleLink {
	s.len = 0
	return s
}

func New() *SingleLink        { return new(SingleLink).Init() }
func (s SingleLink) Len() int { return s.len }

func (s *SingleLink) First() *Element {
	if s.len == 0 {
		return nil
	}
	return s.head
}

// Appends adds an element to the back of a list
func (s *SingleLink) Append(e *Element) {
	if ok := s.minOneElement(); !ok {
		s.head = e
		s.len++
		return
	}

	curr := s.First()
	for i := s.First(); i != nil; i = i.next {
		curr = i
	}

	curr.next = e
	s.len++
}

// Prepend adds an element to the front of a list
func (s *SingleLink) Prepend(e *Element) {
	if ok := s.minOneElement(); !ok {
		s.head = e
		s.len++
		return
	}

	curr := s.First()
	s.head = e
	e.next = curr
	s.len++
}

// Remove removes an element from the list and returns the removed
// element if it was done successfully. In any other case it returns nil.
func (s *SingleLink) Remove(del *Element) *Element {
	if del == nil {
		return nil
	}
	if ok := s.minOneElement(); !ok {
		return nil
	}

	// if the head is the element to remove
	if h := s.First(); del == h {
		s.head = h.next
		s.len--
		return del
	}

	var curr *Element
	for i := s.First(); i != nil && i != del; i = i.next {
		curr = i
	}

	if curr.next == del {
		curr.next = del.next

		// avoid memory leaks
		del.next = nil

		s.len--
		return del
	}

	return nil
}

// Find finds the index of an element. If it does not exist or
// if the list is empty, it returns -1
func (s SingleLink) Find(e *Element) int {
	if ok := s.minOneElement(); !ok {
		return -1
	}

	pos := 0
	for i := s.First(); i != nil; i = i.Next() {
		if i == e {
			return pos
		}
		pos++
	}

	return -1
}

// RemoveBefore removes the element before the given element `e`
// and returns the removed element. If it does not exist or the
// operation is not performed successfully, we return nil.
func (s *SingleLink) RemoveBefore(e *Element) *Element {
	if ok := s.minOneElement(); !ok {
		return nil
	}

	// there is no element before the first one, bail out
	if s.head == e {
		return nil
	}

	/* Should we want to remove the elmement which happens to
	*  be the head, we just move the pointer one step
	*
	* 1.    [ 0x0 ] -> [ 0x1 ]
	*          ^          ^
	*         head        e
	*
	* 2.    [ 0x1 ]
	*          ^
	*         head
	 */
	if s.head.next == e {
		rm := s.head
		s.head = s.head.next
		s.len--
		return rm
	}

	/* If we do not seek to remove the element before the second item in the
	*  list we introduce `trail`: a trailing pointer to backtrack two steps behind
	*  the iterator i. Should we find the item `e`, we link `trail` to `e` which means
	*  we skip over the item in between both of them.
	* 1.    [ 0x0 ] -> [ 0x1 ] -> [ 0x2 ] -> [ 0x3 ]
	*          ^          ^          ^
	*        trail      before e     e
	*        (head)     (delete)    (i)
	*
	*
	* 2.    [ 0x0 ] -> [ 0x2 ] -> [ 0x3 ]
	 */
	trail := s.First()
	for i := trail.next.next; i != nil; i = i.Next() {
		if i == e {
			rm := trail.next
			trail.next = i
			s.len--
			return rm
		}
		trail = trail.next
	}

	return nil
}

// RemoveAfter removes the element after a given element `e`
// and returns the removed element. If it does not exist or the
// operation is not performed successfully, we return nil.
func (s *SingleLink) RemoveAfter(e *Element) *Element {
	if e == nil {
		return nil
	}

	// there is no element to remove
	if e.next == nil {
		return nil
	}

	if ok := s.minOneElement(); !ok {
		return nil
	}

	var target *Element
	for i := s.First(); i != nil; i = i.Next() {
		if i == e {
			target = i
			break
		}
	}

	if target == nil {
		// we never found our element
		return nil
	}

	/* 1.    [ 0x0 ] ----> [ 0x1 ] ----> [ 0x2 ] ----> [ 0x3 ]
	*                         ^             ^             ^
	*                       target     target.next   target.next.next
	*                                  (to remove)
	*
	*
	* 2.     [ 0x0 ] ----> [ 0x1 ] ----> [ 0x3 ]
	*                        ^             ^
	*                      target     target.next
	*
	 */
	removed := target.next
	target.next = target.next.next

	return removed
}

// minOneElement is a helper method which asserts that the list
// has at least one element
func (s *SingleLink) minOneElement() bool {
	if first := s.head; first == nil {
		return false
	}

	return true
}
