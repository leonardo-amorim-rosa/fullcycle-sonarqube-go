package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func time(a int, b int) int {
	return a * b
}

func masterSum(a int, b int) int {
	return a + b + a
}
