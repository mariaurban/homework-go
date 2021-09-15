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

	if (a == b) && (b == c) {
		fmt.Println("Среднее число :", a)
	} else {
		if (a == b) || (a == c) {
			fmt.Println("Среднее число :", a)
		} else {
			if (b == c) || (b == a) {
				fmt.Println("Среднее число :", b)
			} else {
				if ((a < b) && (b < c)) || ((c < b) && (b < a)) {
					fmt.Println("Среднее число:", b)
				} else {
					if ((b < a) && (a < c)) || ((c < a) && (a < b)) {
						fmt.Println("Среднее число:", a)
					} else {
						fmt.Println("Среднее число:", c)
					}
				}
			}
		}
	}
}
