package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is my Go program.")
	var number int
	fmt.Print("Select a number?")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("OS X.")
	case 2:
		fmt.Println("Linux.")
	case 3:
		fmt.Println("Lol loser.")
	default:
		fmt.Printf("%d.", number)
	}
}
