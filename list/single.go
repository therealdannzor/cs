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

func (s *SingleLink) minOneElement() bool {
	if first := s.head; first == nil {
		return false
	}

	return true
}

// TODO:
// GetPos: get index for node n in list
// RemoveBefore: removes node n before an element in list
// RemoveAfter: removes node n after an element in list
