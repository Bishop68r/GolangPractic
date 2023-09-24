package main

import "fmt"

func main() {

	var summa = getSum(5, 5)

	fmt.Println(summa)
}

func getSum(firstNumber int, secondNumber int) int {

	return firstNumber + secondNumber

}
