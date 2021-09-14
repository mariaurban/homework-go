package main

import "fmt"

func main() {
	var a int
	fmt.Println("Пользователь введи первое число НЕ НОЛЬ:")
	fmt.Scanln(&a)

	for a == 0 {
		fmt.Println("Ах ты плохой человек, я же сказала нули не пихать ")
		fmt.Println("Введи нормальное число")
		fmt.Scanln(&a)
	}

	var b int
	fmt.Println("Пользователь введи второе число:")
	fmt.Scanln(&b)

	var с int
	с = a / b
	fmt.Println("Я тут совершил целочисленное деление. И у меня получилось: ", с)
}
