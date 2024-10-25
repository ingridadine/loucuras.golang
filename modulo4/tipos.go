package main

import "fmt"

func main() {
	//bool (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	//string - sequência de bytes
	fmt.Printf("Type: %T - Value: %v\n", "adine", `adine`)
	fmt.Printf("Type: %T - Value: %v\n", "92", "92")

	//int
	fmt.Printf("Type: %T - Value: %v\n", 92, 92)

	//float (float64/float32)
	fmt.Printf("Type: %T - Value: %v\n", 99.99, 99.99)

}
