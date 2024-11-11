package main

import "fmt"

func main() {
	// somaDosValores := soma(42, 13)
	// fmt.Println(somaDosValores)

	// subtraiValores := subtracao(42, 13)
	// fmt.Println(subtraiValores)

	babado1, babado2, _ := revelaBabado("ui ui ui") //O "_" serve para omitir o print da variável babado3
	fmt.Println(babado1)
	fmt.Println(babado2)
	//fmt.Println(babado3)

	nome, sobrenome := printaNomeCompleto("DINYXIA", "NILO")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

func revelaBabado(babado string) (string, string, string) {
	return babado, babado, babado
}

func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// func soma(x int, y int) int {
// 	return x + y
// }

// func subtracao(x int, y int) int {
// 	return x - y
// }
