package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Пользователь введи число:")
	fmt.Scanln(&number)
	fmt.Printf("%04b\n", number)
}
