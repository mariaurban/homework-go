package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Введи координату x: ")
	fmt.Scanln(&x)
	fmt.Println("Введи координату y: ")
	fmt.Scanln(&y)

	if y > 0 {
		if x > 0 {
			fmt.Println("Вы в верхней правой плоскости")
		} else {
			fmt.Println("Вы в нижней правой плоскости")
			if x == 0 {
				fmt.Println("Вы на оси X")
			}
		}
	} else {
		if x > 0 {
			fmt.Println("Вы в верхней левой плоскости")
		} else {
			fmt.Println("Вы в нижней левой плоскости")
		}
		if y == 0 {
			fmt.Println("Вы на оси Y")
		}
		if y == 0 && x == 0 {
			fmt.Println("Вы в начале координат")
		}
	}
}
