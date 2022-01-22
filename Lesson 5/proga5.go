package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Есть массив отсортированный по возрастанию массив чисел.
// Дано число A.
// Найти место для A в массиве, вставить в массив
// Вывести результат

func main() {
	massiv := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}
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
	var massivString []string
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		massivString = append(massivString, strconv.Itoa(massiv[i]))
	}
	fmt.Println(strings.Join(massivString, ","))
}
