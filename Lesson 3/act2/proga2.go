package main

import "fmt"

func main() {
	var a float32
	fmt.Println("Введи число: ")
	fmt.Scanln(&a)
	if a > 0 {
		fmt.Println("Это квадрат числа", a, ": ", a*a)
	} else {
		if a < 0 {
			fmt.Println("Это половина числа", a, ": ", a*0.5)
		} else {
			fmt.Println("Нахуя ты 0 написал ?")
		}
	}

}
