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
			log.Printf("Found %v at %v\n", v, where)
		} else {
			log.Printf("Not Found: %v\n", v)
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
			log.Printf("Found %v at %v\n", v, where)
		} else {
			log.Printf("Not Found: %v\n", v)
		}
	}
}
