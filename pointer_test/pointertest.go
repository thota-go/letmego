package main

import "fmt"

func main() {
	var a = 10
	var b *int
	b = &a
	fmt.Println("Value of a: ", a)
	fmt.Println("Address of a: ", b)
	fmt.Println("Value of a using b: ", *b)

}
