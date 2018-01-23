package list

import (
	"fmt"
)

// Item define Item
type Item struct {
	Value interface{}
	next  *Item
}

func (i *Item) String() string {
	return fmt.Sprintf("%v of type <%T>", i.Value, i.Value)
}

// Get returns value of Item
func (i *Item) Get() interface{} {
	return i.Value
}
