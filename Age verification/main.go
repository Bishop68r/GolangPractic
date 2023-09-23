package main

import "fmt"

func main() {
	var age int8
	fmt.Println("Ваш возраст?")
	fmt.Scan(&age)

	if age <= 18 {
		fmt.Println("Вам нельзя купить алгоголь и пить ведь он с 18 лет, и водить машину")
	} else {
		fmt.Println("Вам разрешается покупать алкоголь и водить тачку")
	}
}
