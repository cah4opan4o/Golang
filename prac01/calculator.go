package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Запрос ввода первого числа
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	// Запрос ввода второго числа
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	// Запрос ввода оператора
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Выполнение операции в зависимости от выбора оператора
	switch operator {
	case "+":
		fmt.Printf("Сумма %.2f и %.2f равна %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Разность %.2f и %.2f равна %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Произведение %.2f и %.2f равно %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Частное %.2f и %.2f равно %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Деление на ноль недопустимо")
		}
	default:
		fmt.Println("Неправильный оператор")
	}
}
