package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)
}

func soma(x int, y int) int {
	return x + y
}
