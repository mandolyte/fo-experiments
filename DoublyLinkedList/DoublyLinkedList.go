package main

import "fmt"

type Node__int struct {
	data int
	next *Node__int
	prev *Node__int
}

type LinkedList__int struct {
	head *Node__int
	tail *Node__int
}

func (list *LinkedList__int) GetFirst() (*Node__int, bool) {
	if list.head == nil {
		return list.head, false
	}
	return list.head, true
}

func (list *LinkedList__int) InsertFirst(i int) {
	data := &Node__int{data: i}
	if list.tail == nil {
		list.head = data
		list.tail = data
		return
	}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList__int) InsertLast(i int) {
	data := &Node__int{data: i}
	if list.head == nil {
		list.head = data
		list.tail = data
		return
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

func (list *LinkedList__int) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}

	return count
}

func (list *LinkedList__int) GetItemsFromStart() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func main() {
	lli := &LinkedList__int{head: nil, tail: nil}
	lli.InsertFirst(1)
	lli.InsertLast(11)
	fmt.Printf("Size of list is:%v\n", lli.GetSize())
	items := lli.GetItemsFromStart()
	for i := range items {
		fmt.Printf("Item %v=%v\n", i, items[i])
	}
}
