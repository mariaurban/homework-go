package main

import "fmt"

func main() {
	var year int
	fmt.Println("Пользователь введи номер года плееез:")
	fmt.Scanln(&year)
	fmt.Print(year+10, " н.э.")

}
