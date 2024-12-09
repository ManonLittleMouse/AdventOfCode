package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func SafeRow(row []int) bool {
	var is_safe = true
	var l_aux []int
	l_aux = append(l_aux, 0)
	for i := 1; i < len(row); i++ {
		l_aux = append(l_aux, row[i-1]-row[i])
	}
	////fmt.Println("\t", l_aux)
	for i := 1; i < len(l_aux); i++ {

		var diff = l_aux[i]
		///fmt.Println(diff, " ")
		if diff == 0 || math.Abs(float64(diff)) > 3 || (l_aux[i]*l_aux[i-1] < 0) {
			is_safe = false

		}
	}
	return is_safe
}

func Dampener(row []int) bool {
	var is_safe = false
	for i := 0; i < len(row); i++ {
		var l_aux []int
		for j := 0; j < len(row); j++ {
			if i != j {
				l_aux = append(l_aux, row[j])
			}
		}
		b := SafeRow(l_aux)
		if b == true {
			is_safe = true
		}
	}
	return is_safe
}

func main() {

	var rows, cols int
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(file), "\n")
	rows = len(lines)

	count := 0

	for i := 0; i < rows-1; i++ {
		line := strings.Split(lines[i], " ")
		cols = len(line)
		var input []int
		for j := 0; j < cols; j++ {
			el, _ := strconv.Atoi(line[j])
			input = append(input, el)
		}
		/// Part one
		///var safe = SafeRow(input)
		/// Part two
		var safe = Dampener(input)
		fmt.Println(input, "---", safe)

		if safe == true {
			count += 1
		}

	}
	fmt.Println(rows)
	fmt.Println(count)

}
