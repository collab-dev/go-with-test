package main

import "fmt"

func main() {
	fmt.Println(add(2, 2))
}

func add(a, b int) int {
	return a * b
}
