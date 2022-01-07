package main

import (
	"fmt"
	"sort"
)

// Есть массив отсортированный по возрастанию массив чисел.
// Дано число A.
// Найти место для A в массиве, вставить в массив
// Вывести результат

func main() {
	massiv := []int{1, 234, 4, 1, 4, 8, 9, 0, 34, 10}
	var a int
	fmt.Println("Наш массив сейчас выглядит так :")
	blyatNakhuyaEtoDelatNuBlin(massiv)
	fmt.Println("Введите любое число :")
	fmt.Scanln(&a)
	massiv = append(massiv, a)
	sort.Ints(massiv)
	blyatNakhuyaEtoDelatNuBlin(massiv)

}

func blyatNakhuyaEtoDelatNuBlin(massiv []int) {
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		fmt.Println(massiv[i])
	}
}
