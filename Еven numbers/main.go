package main

import "fmt"

func main() {
	var input int

	fmt.Println("Введите ваше число:")
	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("Ваш число четное")
	} else {
		fmt.Println("Ваша число нечетное ")
	}

}
