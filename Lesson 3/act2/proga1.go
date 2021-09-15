package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введи число, я его проверять буду: ")
	fmt.Scanln(&a)
	fmt.Println("Число четное: ", 0 == (a%2))
	fmt.Println("Число нечетное: ", 1 == (a%2))
	fmt.Println("Число имеет 2 цифры: ", a > 10)
	fmt.Println("Число имеет 3 цифры: ", a > 100)
	fmt.Println("Число отрицательное: ", a < 0)
	fmt.Println("Число между -10 и 10: ", (a < 10) && (a > -10))
}
