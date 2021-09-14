package main

import (
	"fmt"
)

func main() {
	var asciiNum int
	fmt.Println("Здарова, человек. Введи число ПОЛОЖИТЕЛЬНОЕ ЦЕЛОЕ ЧИСЛО:")
	fmt.Scanln(&asciiNum)
	for asciiNum < 0 {
		fmt.Println("Ах ты плохой человек, это не положительное число5")
		fmt.Println("Введи число ПОЛОЖИТЕЛЬНОЕ ЦЕЛОЕ ЧИСЛО")
		fmt.Println("Желательно более 33:")
		fmt.Scanln(&asciiNum)
	}

	character := string(asciiNum)
	fmt.Println(asciiNum, " : ", character)

}
