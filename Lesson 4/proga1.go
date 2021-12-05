package main

//2. нам ввели число N. вывести все числа от N/2 до N включительно
import "fmt"

func main() {
	fmt.Println("Человек, введи число")
	var N int
	fmt.Scanln(&N)
	fmt.Println("#######Решение#######")
	for i := N / 2; i <= N; i++ {
		fmt.Println(i)
	}
	fmt.Println("##############")

}
