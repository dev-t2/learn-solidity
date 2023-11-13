package main

import "fmt"

var variable1 = "Package Variable"

func main() {
	variable2 := "Function Variable"

	if true {
		variable3 := "Conditional Variable"

		fmt.Println(variable1)
		fmt.Println(variable2)
		fmt.Println(variable3)
	}

	fmt.Println()

	fmt.Println(variable1)
	fmt.Println(variable2)
	// fmt.Println(variable3)

	fmt.Println()

	// num := 5

	// if num > 1 {
	// 	fmt.Println(num)
	// }

	// fmt.Println(num)

	if num := 5; num > 1 {
		fmt.Println(num)
	}

	// fmt.Println(num)
}