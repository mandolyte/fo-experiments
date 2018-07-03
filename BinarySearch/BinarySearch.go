package main

import (
	"log"
	"math"
)

type (
	Comparators__Point struct {
		equals   func(Point, Point) bool
		lessthan func(Point, Point) bool
	}
	Comparators__int struct {
		equals   func(int, int) bool
		lessthan func(int, int) bool
	}
)

func (bs *Comparators__Point) binarySearch(array []Point, val Point) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]

		if bs.equals(val, midItem) {
			return midIndex
		}

		if bs.lessthan(midItem, val) {
			minIndex = midIndex + 1
		} else {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
func (bs *Comparators__int) binarySearch(array []int, val int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]

		if bs.equals(val, midItem) {
			return midIndex
		}

		if bs.lessthan(midItem, val) {
			minIndex = midIndex + 1
		} else {
			maxIndex = midIndex - 1
		}
	}
	return -1
}

func (bs *Comparators__Point) linearSearch(array []Point, val Point) int {
	for index, value := range array {
		if bs.equals(value, val) {
			return index
		}
	}
	return -1
}
func (bs *Comparators__int) linearSearch(array []int, val int) int {
	for index, value := range array {
		if bs.equals(value, val) {
			return index
		}
	}
	return -1
}

func (bs *Comparators__Point) ternarySearch(array []Point, val Point) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex1 := minIndex + int((maxIndex-minIndex)/3)
		midIndex2 := maxIndex - int((maxIndex-minIndex)/3)
		midItem1 := array[midIndex1]
		midItem2 := array[midIndex2]
		if bs.equals(midItem1, val) {
			return midIndex1
		} else if bs.equals(midItem2, val) {
			return midIndex2
		}
		if bs.lessthan(midItem1, val) {
			minIndex = midIndex1 + 1
		} else if bs.lessthan(val, midItem2) {
			maxIndex = midIndex2 - 1
		} else {
			minIndex = midIndex1 + 1
			maxIndex = midIndex2 - 1
		}
	}
	return -1
}
func (bs *Comparators__int) ternarySearch(array []int, val int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex1 := minIndex + int((maxIndex-minIndex)/3)
		midIndex2 := maxIndex - int((maxIndex-minIndex)/3)
		midItem1 := array[midIndex1]
		midItem2 := array[midIndex2]
		if bs.equals(midItem1, val) {
			return midIndex1
		} else if bs.equals(midItem2, val) {
			return midIndex2
		}
		if bs.lessthan(midItem1, val) {
			minIndex = midIndex1 + 1
		} else if bs.lessthan(val, midItem2) {
			maxIndex = midIndex2 - 1
		} else {
			minIndex = midIndex1 + 1
			maxIndex = midIndex2 - 1
		}
	}
	return -1
}

func main() {
	do_int()
	do_point()
}

func do_int() {
	intEq := func(x, y int) bool { return x == y }
	intLt := func(x, y int) bool { return x < y }
	find := &Comparators__int{equals: intEq, lessthan: intLt}

	intarray := []int{10, 20, 30}

	for _, v := range intarray {
		if where := find.binarySearch(intarray, v); where != -1 {
			log.Printf("Binary Found %v at %v\n", v, where)
		} else {
			log.Printf("Binary Not Found: %v\n", v)
		}
		if where := find.linearSearch(intarray, v); where != -1 {
			log.Printf("Linear Found %v at %v\n", v, where)
		} else {
			log.Printf("Linear Not Found: %v\n", v)
		}
		if where := find.ternarySearch(intarray, v); where != -1 {
			log.Printf("Ternary Found %v at %v\n", v, where)
		} else {
			log.Printf("Ternary Not Found: %v\n", v)
		}
	}
}

type Point struct {
	x, y float64
}

func do_point() {
	pointEq := func(a, b Point) bool { return a.x == b.x && a.y == b.y }
	pointLt := func(a, b Point) bool {
		adistance := math.Sqrt((a.x * a.x) + (a.y * a.y))
		bdistance := math.Sqrt((b.x * b.x) + (b.y * b.y))
		return adistance < bdistance
	}
	find := &Comparators__Point{equals: pointEq, lessthan: pointLt}

	pointArray := []Point{{10.0, 10.0}, {20.0, 20.0}, {30.0, 30.0}}

	for _, v := range pointArray {
		if where := find.binarySearch(pointArray, v); where != -1 {
			log.Printf("Binary Found %v at %v\n", v, where)
		} else {
			log.Printf("Binary Not Found: %v\n", v)
		}
		if where := find.linearSearch(pointArray, v); where != -1 {
			log.Printf("Linear Found %v at %v\n", v, where)
		} else {
			log.Printf("Linear Not Found: %v\n", v)
		}
		if where := find.ternarySearch(pointArray, v); where != -1 {
			log.Printf("Ternary Found %v at %v\n", v, where)
		} else {
			log.Printf("Ternary Not Found: %v\n", v)
		}
	}
}
