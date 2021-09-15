package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var between int

	fmt.Println("Введи первое число: ")
	fmt.Scanln(&a)
	fmt.Println("Введи второе число: ")
	fmt.Scanln(&b)
	fmt.Println("Введи третье число: ")
	fmt.Scanln(&c)

	if (a > b) && (a < b) {
		if c > b {
			between = c
		} else {
			between = b
		}
	} else if (b > a) && (b > c) {
		if a > c {
			between = a
		} else {
			between = c
		}
	} else if (c > a) && (c > b) {
		if a > b {
			between = a
		} else {
			between = b
		}
	}
	fmt.Println("Среднее число:", between)
}
