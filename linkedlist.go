package linkedlist

type entry struct {
	val   interface{}
	prev  *entry
	next  *entry
}

type linkedList struct {
	head *entry
	tail *entry
}

func (t linkedList) isEmpty() bool {
	return t.head == nil
}

func (t *linkedList) pushBack(p *entry) {
	if t.head == nil {
		t.head, t.tail = p, p
	} else {
		p.prev = t.tail
		t.tail.next = p
		t.tail = p
	}
}

func (t *linkedList) remove(p *entry) {
	if t.head == p {
		if t.tail == p {
			t.head, t.tail = nil, nil
		} else {
			newHead := t.head.next
			newHead.prev, p.next = nil, nil
			t.head = newHead
		}
	} else if t.tail == p {
		newTail := t.tail.prev
		newTail.next, p.prev = nil, nil
		t.tail = newTail
	} else {
		prev, next := p.prev, p.next
		prev.next, next.prev = next, prev
		p.next, p.prev = nil, nil
	}
}
