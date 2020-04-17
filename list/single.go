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

func (s *SingleLink) Append(e *Element) {
	if ok := s.minOneElement(); !ok {
		s.head = e
		s.len++
		return
	}

	currElement := s.First()
	for i := s.First(); i != nil; i = i.Next() {
		currElement = i
	}

	currElement.next = e
	s.len++

	return
}

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

	return
}

func (s *SingleLink) Remove(del *Element) {
	if ok := s.minOneElement(); !ok {
		return
	}

	var currElement *Element
	for i := s.First(); i != nil && i != del; i = i.Next() {
		currElement = i
	}

	// if the head is the element to remove
	if h := s.First(); del == h {
		s.head = h.next
		return
	}

	if currElement.Next() == del {
		currElement.next = del.next

		// avoid memory leaks
		del.next = nil

		s.len--
	}

	return
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
			return rm
		}
		trail = trail.next
	}

	return nil
}

func (s *SingleLink) minOneElement() bool {
	if first := s.head; first == nil {
		return false
	}

	return true
}

// TODO:
// RemoveAfter: removes node n after an element in list
