package main

import "fmt"

func main() {
	var temp = 33

	if temp > 28 {
		fmt.Println("에어컨을 켠다.")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다.")
	} else {
		fmt.Println("대기한다.")
	}
}