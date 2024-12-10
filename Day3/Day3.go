package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count = 0
	var do = true
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	///test := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	///test2 := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	data := strings.Split("x"+string(file), "mul(")
	for d := 1; d < len(data); d++ {
		mul := strings.Split(data[d], ")")
		expr := strings.Split(mul[0], ",")
		if len(expr) == 2 {
			a := expr[0]
			b := expr[1]
			if !strings.Contains(a, " ") || !strings.Contains(b, " ") {
				u, _ := strconv.Atoi(a)
				v, _ := strconv.Atoi(b)
				/// Part two
				look_for_do := strings.Split(data[d-1], "do")
				last := strings.Split(look_for_do[len(look_for_do)-1], ")")
				if last[0] == "(" {
					do = true
				}
				if last[0] == "n't(" {
					do = false
				}
				if do {
					count += u * v
				}
			}
		}
	}

	fmt.Println(count)

}
