package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberDvo string
	fmt.Println("Пользователь введи десятичное число:")
	fmt.Scanln(&numberDvo)
	number, err := strconv.ParseUint(numberDvo, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(number)

}
