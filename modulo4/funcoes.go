package main

import "fmt"

func main() {
	// somaDosValores := soma(42, 13)
	// fmt.Println(somaDosValores)

	subtraiValores := subtracao(42, 13)
	fmt.Println(subtraiValores)
}

// func soma(x int, y int) int {
// 	return x + y
// }

func subtracao(x int, y int) int {
	return x - y
}
