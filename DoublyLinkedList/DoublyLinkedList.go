package main

import "log"

type (
	Node__float64 struct {
		data float64
		next *Node__float64
		prev *Node__float64
	}
	Node__int struct {
		data int
		next *Node__int
		prev *Node__int
	}
	Node__string struct {
		data string
		next *Node__string
		prev *Node__string
	}
)

func (n *Node__float64) Value() float64 {
	return n.data
}
func (n *Node__int) Value() int {
	return n.data
}
func (n *Node__string) Value() string {
	return n.data
}

type (
	LinkedList__float64 struct {
		head *Node__float64
		tail *Node__float64
	}
	LinkedList__int struct {
		head *Node__int
		tail *Node__int
	}
	LinkedList__string struct {
		head *Node__string
		tail *Node__string
	}
)

func (list *LinkedList__float64) GetFirst() (*Node__float64, bool) {
	if list.head == nil {
		return list.head, false
	}
	return list.head, true
}
func (list *LinkedList__int) GetFirst() (*Node__int, bool) {
	if list.head == nil {
		return list.head, false
	}
	return list.head, true
}
func (list *LinkedList__string) GetFirst() (*Node__string, bool) {
	if list.head == nil {
		return list.head, false
	}
	return list.head, true
}

func (list *LinkedList__float64) GetLast() (*Node__float64, bool) {
	if list.head == nil {
		return list.head, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current, true
}
func (list *LinkedList__int) GetLast() (*Node__int, bool) {
	if list.head == nil {
		return list.head, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current, true
}
func (list *LinkedList__string) GetLast() (*Node__string, bool) {
	if list.head == nil {
		return list.head, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current, true
}

func (list *LinkedList__float64) InsertFirst(i float64) {
	data := &Node__float64{data: i}
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
func (list *LinkedList__string) InsertFirst(i string) {
	data := &Node__string{data: i}
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

func (list *LinkedList__float64) InsertLast(i float64) {
	data := &Node__float64{data: i}
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
func (list *LinkedList__string) InsertLast(i string) {
	data := &Node__string{data: i}
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

func (list *LinkedList__float64) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}

	return count
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
func (list *LinkedList__string) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}

	return count
}

func (list *LinkedList__float64) GetItemsFromStart() []float64 {
	var items []float64
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
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
func (list *LinkedList__string) GetItemsFromStart() []string {
	var items []string
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *LinkedList__float64) GetItemsFromEnd() []float64 {
	var items []float64
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}
func (list *LinkedList__int) GetItemsFromEnd() []int {
	var items []int
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}
func (list *LinkedList__string) GetItemsFromEnd() []string {
	var items []string
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}

func (list *LinkedList__float64) RemoveByValue(i float64) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}
func (list *LinkedList__int) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}
func (list *LinkedList__string) RemoveByValue(i string) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList__float64) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}
func (list *LinkedList__int) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}
func (list *LinkedList__string) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList__float64) SearchValue(i float64) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}
func (list *LinkedList__int) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}
func (list *LinkedList__string) SearchValue(i string) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	do_int()
	do_float()
	do_string()
}

func do_int() {
	log.Println("\n----- Int Tests -----")
	lli := &LinkedList__int{head: nil, tail: nil}
	lli.InsertFirst(1)
	lli.InsertLast(3)
	lli.InsertLast(5)
	lli.InsertLast(7)
	log.Printf("Size of list is:%v\n", lli.GetSize())
	x, ok := lli.GetFirst()
	if ok {
		log.Printf("Head value is: %v\n", x.Value())
	} else {
		log.Fatal("Head is missing")
	}

	x, ok = lli.GetLast()
	if ok {
		log.Printf("Tail value is: %v\n", x.Value())
	} else {
		log.Fatal("Tail is missing")
	}

	log.Print("All items from head\n")
	items := lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
		if !lli.SearchValue(items[i]) {
			log.Fatalf("Error! value %v not found\n", items[i])
		}
	}

	if lli.SearchValue(-1) {
		log.Fatalf("Error! non-existent value found: %v\n", -1)
	}

	log.Print("All items from tail (backwards)\n")
	items = lli.GetItemsFromEnd()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}

	if ok = lli.RemoveByIndex(1); !ok {
		log.Fatal("Error! lli.RemoveByIndex(1)\n")
	}

	if ok = lli.RemoveByValue(5); !ok {
		log.Fatal("Error! lli.RemoveByValue(5)\n")
	}

	log.Print("Final items after removals (should be 1 and 7)\n")
	items = lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}
}

func do_float() {
	log.Println("\n----- Float Tests -----")

	lli := &LinkedList__float64{head: nil, tail: nil}
	lli.InsertFirst(1.1)
	lli.InsertLast(3.3)
	lli.InsertLast(5.5)
	lli.InsertLast(7.7)
	log.Printf("Size of list is:%v\n", lli.GetSize())
	x, ok := lli.GetFirst()
	if ok {
		log.Printf("Head value is: %v\n", x.Value())
	} else {
		log.Fatal("Head is missing")
	}

	x, ok = lli.GetLast()
	if ok {
		log.Printf("Tail value is: %v\n", x.Value())
	} else {
		log.Fatal("Tail is missing")
	}

	log.Print("All items from head\n")
	items := lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
		if !lli.SearchValue(items[i]) {
			log.Fatalf("Error! value %v not found\n", items[i])
		}
	}

	if lli.SearchValue(-1) {
		log.Fatalf("Error! non-existent value found: %v\n", -1)
	}

	log.Print("All items from tail (backwards)\n")
	items = lli.GetItemsFromEnd()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}

	if ok = lli.RemoveByIndex(1); !ok {
		log.Fatal("Error! lli.RemoveByIndex(1)\n")
	}

	if ok = lli.RemoveByValue(5.5); !ok {
		log.Fatal("Error! lli.RemoveByValue(5.5)\n")
	}

	log.Print("Final items after removals (should be 1.1 and 7.7)\n")
	items = lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}
}

func do_string() {
	log.Println("\n----- String Tests -----")

	lli := &LinkedList__string{head: nil, tail: nil}
	lli.InsertFirst("a")
	lli.InsertLast("b")
	lli.InsertLast("c")
	lli.InsertLast("d")
	log.Printf("Size of list is:%v\n", lli.GetSize())
	x, ok := lli.GetFirst()
	if ok {
		log.Printf("Head value is: %v\n", x.Value())
	} else {
		log.Fatal("Head is missing")
	}

	x, ok = lli.GetLast()
	if ok {
		log.Printf("Tail value is: %v\n", x.Value())
	} else {
		log.Fatal("Tail is missing")
	}

	log.Print("All items from head\n")
	items := lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
		if !lli.SearchValue(items[i]) {
			log.Fatalf("Error! value %v not found\n", items[i])
		}
	}

	if lli.SearchValue("z") {
		log.Fatalf("Error! non-existent value found: %v\n", "z")
	}

	log.Print("All items from tail (backwards)\n")
	items = lli.GetItemsFromEnd()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}

	if ok = lli.RemoveByIndex(1); !ok {
		log.Fatal("Error! lli.RemoveByIndex(1)\n")
	}

	if ok = lli.RemoveByValue("c"); !ok {
		log.Fatal("Error! lli.RemoveByValue(\"c\")\n")
	}

	log.Print("Final items after removals (should be a and d)\n")
	items = lli.GetItemsFromStart()
	for i := range items {
		log.Printf("Item %v=%v\n", i, items[i])
	}
}
