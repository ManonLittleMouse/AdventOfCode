package main

import (
	"fmt"
	"math"
)

func FindSmallest(l []int) (int, []int) {
	// Compute the minimal element
	min := l[0]
	for i := 0; i < len(l); i++ {
		if l[i] < min {
			min = l[i]
		}
	}
	// Compute the rest
	var lf []int
	for i := 0; i < len(l); i++ {
		if l[i] != min {
			lf = append(lf, l[i])
		}
	}
	return min, lf
}

func ComputeDiff(u, v []int) int {
	/// We suppose the two lists have the same size
	if len(u) != len(v) {
		fmt.Println("Size of list not equal!")
		return -1
	}
	var count int
	var done = false
	for done == false {
		el_u, list_u := FindSmallest(u)
		el_v, list_v := FindSmallest(v)
		count += int(math.Abs(float64(el_u) - float64(el_v)))
		if len(list_u) == 0 {
			done = true
		} else {
			u = list_u
			v = list_v
		}
	}
	return count
}

func main() {
	var l1 = []int{3, 4, 2, 1, 3, 3}
	var l2 = []int{4, 3, 5, 3, 9, 3}
	var count = ComputeDiff(l1, l2)
	fmt.Println(count)
}
