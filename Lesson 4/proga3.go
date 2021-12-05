package main

import (
	"fmt"
)

//вывести числа от 1 до 100 по себе раз (1 раз 1, 2 раза 2...). Каждое новое
//число на новой строке
func main() {
	fmt.Println("Человек, введи число от 1 до 100")
	var N int
	fmt.Scanln(&N)
	fmt.Println("#######Решение#######")
	if N <= 100 {
		for i := 1; i < N+1; i++ {
			for j := 1; j < i+1; j++ {
				fmt.Print(i)
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("Ля ты какой, введи число от 1 до 100, а не", N)
	}
}
