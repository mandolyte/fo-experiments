package main

import "fmt"

type Point struct {
	x,y int
}

func main() {
	i := *new(int)
	r := *new(float64)
	s := *new(string)
	b := *new(bool)
	p := *new(Point)
	a := *new([]int)
	m := *new(map[int]string)

	fmt.Printf("Int:%v\n",i)
	fmt.Printf("Float64:%v\n",r)
	fmt.Printf("String:%v\n",s)
	fmt.Printf("Bool:%v\n",b)
	fmt.Printf("Point:%v\n",p)
	fmt.Printf("Slice:%v\n",a)
	fmt.Printf("Map:%v\n",m)
}


