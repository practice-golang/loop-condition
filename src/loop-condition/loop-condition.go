package main

import "fmt"

func main() {

	// if 같은 for
	i := 0
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("----")

	// 그냥 for에 그냥 switch
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Unknown")
		}
	}

	fmt.Println("----")

	// 그냥 for에 if 같은 switch
	for i := 0; i < 5; i++ {
		switch {
		case i == 0:
			fmt.Println("Zero")
		case i == 1:
			fmt.Println("One")
		case i == 2:
			fmt.Println("Two")
		case i > 2:
			fmt.Println("More than two")
		default:
			fmt.Println("Unknown")
		}
	}

	fmt.Println("----")

	// 그냥 if
	i = 10
	if i < 11 {
		fmt.Println(i)
	}

	fmt.Println("----")

	// for 같은 if
	if i := 10; i > 9 {
		fmt.Println(i)
	}
}
