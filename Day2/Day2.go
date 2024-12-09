package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows int
	Cols int
	data []int
}

func (m *Matrix) At(row, col int) int {
	return m.data[row*m.Cols+col]
}

func SafeRow(m *Matrix, r int) bool {
	var is_safe = true
	var l_aux []int
	l_aux = append(l_aux, m.At(r, 0))
	for i := 1; i < m.Cols; i++ {
		l_aux = append(l_aux, l_aux[len(l_aux)-1]-m.At(r, i))
	}
	for i := 1; i < len(l_aux); i++ {
		var diff = int(math.Abs(float64(l_aux[i] - l_aux[i-1])))
		if diff == 0 || diff > 3 {
			is_safe = false
		}
	}
	return is_safe
}

func main() {
	var input = &Matrix{
		Rows: 6,
		Cols: 5,
		data: []int{7, 6, 4, 2, 1, 1, 2, 7, 8, 9, 9, 7, 6, 2, 1, 1, 3, 2, 4, 5, 8, 6, 4, 4, 1, 1, 3, 6, 7, 9},
	}
	var count = 0
	for r := 0; r < input.Rows; r++ {
		var safe = SafeRow(input, r)
		if safe == false {
			count += 1
		}
	}
	fmt.Println(count)

}
