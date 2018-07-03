package main

import "fmt"

type (
	CircularBuffer__float64 struct {
		data    []float64
		pointer int
		size    int
	}
	CircularBuffer__int struct {
		data    []int
		pointer int
		size    int
	}
	CircularBuffer__string struct {
		data    []string
		pointer int
		size    int
	}
)

func (b *CircularBuffer__float64) InsertValue(i float64) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}
func (b *CircularBuffer__int) InsertValue(i int) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}
func (b *CircularBuffer__string) InsertValue(i string) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}

func (b *CircularBuffer__float64) GetValues() []float64 {
	return b.data
}
func (b *CircularBuffer__int) GetValues() []int {
	return b.data
}
func (b *CircularBuffer__string) GetValues() []string {
	return b.data
}

func (b *CircularBuffer__float64) GetValuesFromPosition(i int) ([]float64, bool) {

	if i >= b.size {
		return nil, false
	}

	out := make([]float64, b.size)
	for u := 0; u < b.size; u++ {
		if i >= b.size {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}
func (b *CircularBuffer__int) GetValuesFromPosition(i int) ([]int, bool) {

	if i >= b.size {
		return nil, false
	}

	out := make([]int, b.size)
	for u := 0; u < b.size; u++ {
		if i >= b.size {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}
func (b *CircularBuffer__string) GetValuesFromPosition(i int) ([]string, bool) {

	if i >= b.size {
		return nil, false
	}

	out := make([]string, b.size)
	for u := 0; u < b.size; u++ {
		if i >= b.size {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}

func New__float64(sz int) *CircularBuffer__float64 {
	cdata := make([]float64, sz)
	cb := &CircularBuffer__float64{data: cdata, pointer: 0, size: sz}
	return cb
}
func New__int(sz int) *CircularBuffer__int {
	cdata := make([]int, sz)
	cb := &CircularBuffer__int{data: cdata, pointer: 0, size: sz}
	return cb
}
func New__string(sz int) *CircularBuffer__string {
	cdata := make([]string, sz)
	cb := &CircularBuffer__string{data: cdata, pointer: 0, size: sz}
	return cb
}

func main() {
	do_int()
	do_string()
	do_float()
}

func do_int() {
	cb := New__int(3)
	cb.InsertValue(10)
	cb.InsertValue(20)
	cb.InsertValue(30)
	cb.InsertValue(40)
	cb.InsertValue(50)
	cbslice := cb.GetValues()
	fmt.Println("--- All Items ---")
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}

	fmt.Println("--- Items from end---")
	cbslice, ok := cb.GetValuesFromPosition(2)
	if !ok {
		fmt.Printf("Error: position too large: %v\n", 2)
	}
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}
}

func do_string() {
	cb := New__string(3)
	cb.InsertValue("ten")
	cb.InsertValue("twenty")
	cb.InsertValue("thirty")
	cb.InsertValue("forty")
	cb.InsertValue("fifty")
	cbslice := cb.GetValues()
	fmt.Println("--- All Items ---")
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}

	fmt.Println("--- Items from end---")
	cbslice, ok := cb.GetValuesFromPosition(2)
	if !ok {
		fmt.Printf("Error: position too large: %v\n", 2)
	}
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}
}

func do_float() {
	cb := New__float64(3)
	cb.InsertValue(10.10)
	cb.InsertValue(20.20)
	cb.InsertValue(30.30)
	cb.InsertValue(40.40)
	cb.InsertValue(50.50)
	cbslice := cb.GetValues()
	fmt.Println("--- All Items ---")
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}

	fmt.Println("--- Items from end---")
	cbslice, ok := cb.GetValuesFromPosition(2)
	if !ok {
		fmt.Printf("Error: position too large: %v\n", 2)
	}
	for x := range cbslice {
		fmt.Printf("Item %v = %v\n", x, cbslice[x])
	}
}
