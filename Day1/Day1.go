package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindSmallest(l []int) (int, []int) {
	/// Verify that the lenght is > to 0
	if len(l) <= 0 {
		print("Error, empty list!")
	}
	// Compute the minimal element
	min := l[0]
	for i := 0; i < len(l); i++ {
		if l[i] < min {
			min = l[i]
		}
	}
	// Compute the rest
	var lf []int
	first_see := true /// To handle identical int

	for i := 0; i < len(l); i++ {
		if l[i] != min || first_see == false {
			lf = append(lf, l[i])
		} else {
			first_see = false
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
		print(el_u, " ", el_v, " ", el_u-el_v, "\n")
		if len(list_u) != len(list_v) {
			print("el_u: ", el_u, " el_v: ", el_v, "   size_u ", len(list_u), " size_v ", len(list_v), "\n")
		}
		if len(list_u) == 0 {
			done = true
		} else {
			u = list_u
			v = list_v
		}
	}
	return count
}

func ComputeSim(u, v []int) int {
	count := 0
	for i := 0; i < len(u); i++ {
		el := u[i]
		sim := 0
		for j := 0; j < len(v); j++ {
			if el == v[j] {
				sim++
			}
		}
		count += el * sim
	}
	return count
}
func main() {
	var l1 []int
	var l2 []int
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines)-1; i++ {
		line := strings.Split(lines[i], "   ")
		var s = line[0]
		var i1, _ = strconv.Atoi(s)
		l1 = append(l1, i1)
		var s2 = line[1]
		var i2, _ = strconv.Atoi(s2)
		l2 = append(l2, i2)
	}

	/// Part One

	//var count = ComputeDiff(l1, l2)
	// fmt.Println(count)

	/// Part two

	var count2 = ComputeSim(l1, l2)
	print(count2)

}
