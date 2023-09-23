package main

import (
	"fmt"
)

func main() {
	var FirstNumber int
	var SecondNumber int
	var Selection int
	fmt.Println("Добро пожаловать, в какулятор выражений! ")
	fmt.Println("Введите первое число:")
	fmt.Scan(&FirstNumber)
	fmt.Println("Введите первое число: ")
	fmt.Scan(&SecondNumber)

	fmt.Println("Выбирайте  любую операцию \n 1 Сложение \n 2 Вычитание \n 3 Деление \n 4 Умножение ")
	fmt.Scan(&Selection)

	switch Selection {
	case 1:
		fmt.Println("Результат сложения: ", FirstNumber+SecondNumber)
	case 2:
		fmt.Println("Результат вычитания: ", FirstNumber-SecondNumber)
	case 3:
		fmt.Println(" Результат деления: ", FirstNumber/SecondNumber)
	case 4:
		fmt.Println(" Результат умножения: ", FirstNumber*SecondNumber)
	default:
		fmt.Println("значение  выражения  не определено")

	}

}
