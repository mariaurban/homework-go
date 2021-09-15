package main

import "fmt"

func main() {
	var year int
	fmt.Println("Введи год: ")
	fmt.Scanln(&year)

	if 0 != year%4 || (year%100 == 0 && year%400 != 0) {
		fmt.Println("Обычный")
	} else {
		fmt.Println("Високосный")
	}
}
