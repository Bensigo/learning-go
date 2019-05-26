package main

import "fmt"

// function that return a single value
func add(a, b int) int {
	return a + b
}

func concat(first string, second string) string {
	return first + " " + second
}
func main() {
	fmt.Println(add(33, 10))
	fmt.Println(concat("Bensigo", "Egwey"))
}
