package main

import (
	"fmt"
)

func main() {
	fmt.Println("Здарова, человек. Введи свое имя :")
	var name string
	fmt.Scanln(&name)
	fmt.Println("А теперь свою фамилию :")
	var surname string
	fmt.Scanln(&surname)
	pidor := name + " " + surname
	fmt.Println("Ну здарова, ", pidor)
}
