package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Введи первое число: ")
	fmt.Scanln(&a)
	fmt.Println("Введи второе число: ")
	fmt.Scanln(&b)
	fmt.Println("Введи третье число: ")
	fmt.Scanln(&c)

	if a == 0 && b == 0 && c == 0 {
		fmt.Println("Дурак, зачем ты нули ввел?")
	} else {
		averageValue := (a + b + c) / 2
		fmt.Println("Среднее число: ", averageValue)
	}

}
