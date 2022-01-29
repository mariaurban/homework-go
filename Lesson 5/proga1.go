package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Есть массив чисел.
// Найти в них самое максимальное и минимальное число - поменять их местами.
// Вывести результирующий массив.

func main() {
	var numb []int
	var numb1 int
	fmt.Println("Введи пожалуйста элементы массива, 5 шт")
	for i := 0; i < 5; i++ {
		fmt.Scanln(&numb1)
		numb = append(numb, numb1)
	}
	fmt.Println("Наш массив выглятит так")
	blyatNakhuyaEtoDelat(numb)

	mini := 0
	maxi := 0

	for i, v := range numb {
		if v < i {
			mini = i
		}
		if v > i {
			maxi = i
		}
	}
	fmt.Println("Максимальное число в массиве :", numb[maxi])
	fmt.Println("Минимальное число в массиве :", numb[mini])
	fmt.Println("Сейчас я поменяю местами минимальное и максимальное число ")

	max := numb[mini]
	numb[mini] = numb[maxi]
	numb[maxi] = max

	fmt.Println("Теперь наш массив выгляди так")
	blyatNakhuyaEtoDelat(numb)
}
func blyatNakhuyaEtoDelat(massiv []int) {

	var massivString []string
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		massivString = append(massivString, strconv.Itoa(massiv[i]))
	}
	fmt.Println(strings.Join(massivString, ","))
}
