package main

import "fmt"

func add(x, y int) int {
	return x + y
}

var c, python, java = true, false, "no!"

func swap(x, y string) (string, string) {
	return x + y, y
}

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
