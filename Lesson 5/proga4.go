package main

import "fmt"

// Есть массив чисел, есть число меньшее чем размер массива.
// Удалить элемент с этим номером из массива.
// Вывести результирующий масиив.

func main() {
	number := []int{1, 23, 12, 2, 4, 4, 2, 19, 12, 9}
	var i int
	fmt.Println(" Наш массив сейчас выглядит так :", number)
	fmt.Println("Введите число, до 10: ")
	fmt.Scanln(&i)
	number = append(number[:i-1], number[i:]...)
	print("Я удалила элемент под номером", i, "из своего массива")
	blyatNakhuyaEtoDelatNuRili(number)
}
func blyatNakhuyaEtoDelatNuRili(massiv []int) {
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		fmt.Println(massiv[i])
	}
}
