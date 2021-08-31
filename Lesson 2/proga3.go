package main

import (
	"fmt"
)

func main() {
	fmt.Println("Человек, введи число")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Человек, введи еще одно число")
	var num2 int
	fmt.Scanln(&num2)
	sum := num1 + num2
	dif := num1 - num2
	mult := num1 * num2
	var delim float64
	delim = float64(num1 / num2)
	ost := num1 % num2
	fmt.Println("Таки я сложила ваши два числа и получилось: ", sum)
	fmt.Println("Таки я из первого числа вычла второе и получилось: ", dif)
	fmt.Println("Таки я умножила ваши два числа и получилось: ", mult)
	fmt.Println("Таки я разделила первое число на второе получилось: ", delim)
	fmt.Println("Таки я разделил первое число на второе и остаток получился: ", ost)
}
