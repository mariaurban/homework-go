package main

import (
	"fmt"
)

func main() {
	var num string
	fmt.Print("Чтобы выйти на связь с оператором нашмите 6 ")
	fmt.Scanln(&num)
	switch num {
	case "6":
		fmt.Println("Переводим вас на оператора. Ожидайте. Ваш номер в очереди 1488")
	default:
		fmt.Printf("Не правильное число")
	}
}
