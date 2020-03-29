package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is my Go program.")
	var number int
	var value int
	var num int
	num = 0
	value = 1
	for value <= 2 {
		fmt.Println("\n\nPlease select an option:")
		fmt.Println("1 - Print Covid19 cases in Pakistan")
		fmt.Println("2 - Print Covid19 cases in South Korea")
		fmt.Println("3 - Print Covid19 cases in France")
		fmt.Println("4 - Print my personalized message to Corona Virus!")
		fmt.Println("0 - Exit")
		fmt.Scanf("%d", &number)
		fmt.Println()
		switch number {
		case 0:
			{
				fmt.Println("Exiting. Goodbye.")
				if num >= 10 {
					value = 3
					break
				} else {
					fmt.Println("View all the country options first :) ")
				}
			}
		case 1:
			fmt.Println("Option 1")
			fmt.Println("Pakistan has 1560 cases so far.")
			num++
		case 2:
			fmt.Println("Option 2")
			fmt.Println("South Korea has 9583 cases so far.")
			num = num + 2
		case 3:
			fmt.Println("Option 3")
			fmt.Println("France has 37,611 cases so far.")
			num = num + 3
		case 4:
			fmt.Println("Option 4")
			fmt.Println("Corona se darhna nai,\nLarna hai :) 👊")
			num = num + 4
		default:
			fmt.Println("Invalid selection. Try again.")
		}
	}

	/* 	fmt.Print("Select a number?")

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
	   	sum := 1
	   	for sum <= 10 {
	   		fmt.Println(sum)
	   		sum++
	   	}

	   	for i := 0; i < 5; i++ {
	   		sum += i
	   		fmt.Printf("%d ", sum)
	   	}
	   	fmt.Println()
	   	fmt.Println(sum) */

}
