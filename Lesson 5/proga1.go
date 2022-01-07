package main

import "fmt"

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
	min := numb[0]
	max := numb[0]
	mini := 0
	maxi := 0
	for i, v := range numb {
		if v < min {
			min = v
			mini = i
		}
		if v > max {
			max = v
			maxi = i
		}
	}
	fmt.Println("Максимальное число в массиве :", max)
	fmt.Println("Минимальное число в массиве :", min)
	fmt.Println("Сейчас я поменяю местами минимальное и максимальное число ")

	numb[mini] = max
	numb[maxi] = min

	fmt.Println("Теперь наш массив выгляди так", numb)

}
func blyatNakhuyaEtoDelat(massiv []int) {
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		fmt.Println(massiv[i])
	}

}
