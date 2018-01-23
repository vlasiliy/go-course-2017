package list

import (
	"errors"
)

// List structure that define List
type List struct {
	first  *Item
	cursor *Item
}

// New create new list
func New() *List {
	return new(List)
}

// Next return next element
func (l *List) Next() (item *Item, err error) {
	if l.first == nil {
		return nil, errors.New("List is empty")
	}

	if l.cursor == nil {
		return nil, errors.New("End reached")
	}

	next := l.cursor
	l.cursor = l.cursor.next

	return next, nil
}

// Empty return true if list is empty
func (l *List) Empty() bool {
	if l.first == nil {

		return true
	}

	return false
}

// Len return length of list
func (l *List) Len() (count uint) {
	for i := l.first; i != nil; i = i.next {
		count++
	}

	return
}

// Unshift add item to the beginning of the list
func (l *List) Unshift(value interface{}) error {
	defer l.resetCursor()

	i := Item{value, nil}

	if l.first != nil {
		i.next = l.first
	}

	l.first = &i

	return nil
}

// Insert add item after N skipped elements
func (l *List) Insert(value interface{}, index uint) error {
	defer l.resetCursor()

	item := l.first

	var i uint

	for {
		if index == i {
			next := item.next
			item.next = &Item{value, next}

			return nil
		}

		if item.next == nil {
			return errors.New("Out of range")
		}

		i++
		item = item.next
	}
}

// Append add element at the end of list
func (l *List) Append(value interface{}) error {
	defer l.resetCursor()

	item := l.first

	for {
		if item.next == nil {
			item.next = &Item{value, nil}

			return nil
		}

		item = item.next
	}
}

// Delete element by the index
func (l *List) Delete(index uint) error {
	defer l.resetCursor()

	var i uint
	var prev = l.first

	if index == 0 {
		l.first = (*l.first).next
	}

	for {
		i++

		if index == i {
			prev.next = (*prev.next).next

			return nil
		}

		if prev.next == nil {
			return errors.New("Out of range")
		}

		prev = prev.next
	}
}

func (l *List) resetCursor() {
	l.cursor = l.first
}

func (l *List) isCursorPointToEnd() bool {
	if l.cursor == nil {
		return true
	}

	return false
}
