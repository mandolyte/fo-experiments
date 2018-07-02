package main

import "fmt"

func reverseArray__float64(a []float64) {
	i := 0
	u := len(a) - 1
	for i < u {
		a[i], a[u] = a[u], a[i]
		i, u = i+1, u-1
	}
}
func reverseArray__int(a []int) {
	i := 0
	u := len(a) - 1
	for i < u {
		a[i], a[u] = a[u], a[i]
		i, u = i+1, u-1
	}
}
func reverseArray__string(a []string) {
	i := 0
	u := len(a) - 1
	for i < u {
		a[i], a[u] = a[u], a[i]
		i, u = i+1, u-1
	}
}

func main() {
	ia := []int{1, 2, 3}
	reverseArray__int(ia)
	fmt.Println(ia)

	sa := []string{"a", "b", "c"}
	reverseArray__string(sa)
	fmt.Println(sa)

	ra := []float64{1.1, 2.2, 3.3}
	reverseArray__float64(ra)
	fmt.Println(ra)
}
