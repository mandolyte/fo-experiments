package main

import "fmt"

func bubbleSort__float64(gt func(float64, float64) bool, array []float64) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {

			if gt(array[itemIndex], array[itemIndex+1]) {
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}
func bubbleSort__int(gt func(int, int) bool, array []int) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {

			if gt(array[itemIndex], array[itemIndex+1]) {
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}
func bubbleSort__string(gt func(string, string) bool, array []string) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {

			if gt(array[itemIndex], array[itemIndex+1]) {
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}

func main() {
	ia := []int{4, 2, 3}
	gti := func(x, y int) bool { return x > y }
	bubbleSort__int(gti, ia)
	fmt.Println(ia)

	sa := []string{"z", "b", "c"}
	gts := func(x, y string) bool { return x > y }
	bubbleSort__string(gts, sa)
	fmt.Println(sa)

	ra := []float64{4.4, 2.2, 3.3}
	gtf := func(x, y float64) bool { return x > y }
	bubbleSort__float64(gtf, ra)
	fmt.Println(ra)

}
