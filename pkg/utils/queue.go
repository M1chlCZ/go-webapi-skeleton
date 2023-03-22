package utils

import (
	"container/list"
	"errors"
)

type Queue struct {
	items *list.List
}

func (c *Queue) Enqueue(value string) {
	c.items.PushBack(value)
}

func (c *Queue) Dequeue() {
	if c.items.Len() > 0 {
		element := c.items.Front()
		c.items.Remove(element)
	}
}

func (c *Queue) Front() (*list.Element, error) {
	if c.items.Len() > 0 {
		return c.items.Front(), nil
	}
	return nil, errors.New("queue is empty")
}

func (c *Queue) Size() int {
	return c.items.Len()
}

//// Initialize queue
//queue := &Queue{
//items: list.New(),
//}
//// Add items to queue
//queue.Enqueue("test 1")
//queue.Enqueue("test 2")
//fmt.Println("Size: " + strconv.Itoa(queue.Size()))
//// Process items until size > 0
//for queue.Size() > 0 {
//// Get next value
//value, err := queue.Front()
//if err == nil {
//fmt.Printf("Value: %s\n", value.Value)
//// Remove item from queue
//queue.Dequeue()
//}
//}
//// Output:
////Size: 2
////Value: test 1
////Value: test 2
////Size: 0
//fmt.Printf("Size: %d\n", queue.Size())
