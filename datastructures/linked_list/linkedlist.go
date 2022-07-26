package linkedlist

import "fmt"

type node struct {
	value interface{}
	next  *node
  prev  *node
}

type LinkedList struct {
	head   *node
	tail   *node
	Length int
}

func New() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) Append(val interface{}) {
	newNode := node{value: val, prev: ll.tail}
	if ll.head == nil {
		ll.head = &newNode
		ll.tail = ll.head
	} else {
		ll.tail.next = &newNode
		ll.tail = &newNode
	}
	ll.Length += 1
	return
}

func (ll *LinkedList) Prepend(val interface{}) {
	newNode := node{value: val, next: ll.head}
  ll.head.prev = &newNode
	ll.head = &newNode
	ll.Length += 1
}

func (ll *LinkedList) Insert(position int, val interface{}) error {
	if position < 0 {
		return fmt.Errorf("position cannot be negative")
	}

	if position >= ll.Length {
		ll.Append(val)
		ll.Length += 1
		return nil
	}

	if position == 0 {
		ll.Prepend(val)
		ll.Length += 1
		return nil
	}

	before := ll.traverseToPosition(position - 1)
	newNode := node{value: val, next: before.next, prev: before}
	before.next = &newNode
  newNode.next.prev = &newNode
  ll.Length++

	return nil
}

func (ll *LinkedList) Remove(position int) error {
	if position < 0 || position > ll.Length {
		return fmt.Errorf("position %d out of range", position)
	}

	before := ll.traverseToPosition(position - 1)
	before.next = before.next.next
  before.next.prev = before
	ll.Length--
	return nil
}

func (ll LinkedList) traverseToPosition(position int) *node {
	var i int
	curr := ll.head
	for i != position {
		curr = curr.next
		i++
	}
	return curr
}

func (ll *LinkedList) Print() {
	values := []interface{}{}
	curr := ll.head
	for curr != nil {
		values = append(values, curr.value)
		curr = curr.next
	}
	fmt.Println(values)
}

func (ll *LinkedList) Reverse() {
  if ll.head.next == nil {
    return
  }

  first := ll.head
  ll.tail = ll.head
  second := first.next

  for second != nil {
    tmp := second.next
    second.next = first
    first = second
    second = tmp
  }

  ll.head.next = nil
  ll.head = first
}