package main

import "fmt"

//дано число N от 1 до 52. Вывести первые N символов латинского
//алфавита. Сначала заглваные. Если заглавные закончились - выводим прописные.
//Например:
//3 -> ABC
//30 -> ABC...XYZabсd
func main() {
	fmt.Println("Человек, введи число ОТ 1 ДО 52 ")
	var N int
	fmt.Scanln(&N)
	fmt.Println("#######Решение#######")
	if N <= 52 {
		for i := 65; i < 65+N; i++ {
			if i > 90 {
				fmt.Print(string(i + 6))
			} else {
				fmt.Print(string(i))
			}
		}
	} else {
		fmt.Println("Человек, введи число ОТ 1 ДО 52, а ты дурак и ввел :", N)
	}
}
