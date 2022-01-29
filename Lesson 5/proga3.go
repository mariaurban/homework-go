package main

import (
	"fmt"
	"strings"
)

// Алфавит: Есть два массива символов.
// Собрать третий массив из символов которые есть и в том и в другом массиве.
// Каждый символ в результате не упоминается больше одного раза
// Вывести третий массив

func main() {
	array1 := [12]rune{'2', '4', '2', 'f', 'q', 'e', '2', 's', 'D', 'w', '6', '3'}
	array2 := [12]rune{'3', '2', '7', '6', 'd', 'a', '3', 'e', 'w', 'r', 'p', 'l'}
	var alfa []rune
	fmt.Println("Первый массив: ")
	blyatNakhuyaEtoDelatSukaSuka12(array1)
	fmt.Println("Второй массив: ")
	blyatNakhuyaEtoDelatSukaSuka12(array2)
	//for _, v := range array1 {
	//	if !contains(alfa, v) {
	//		alfa = append(alfa, v)
	//	}
	//}
	//for _, v := range array2 {
	//	if !contains(alfa, v) {
	//		alfa = append(alfa, v)
	//	}
	//}
	for _, v := range array1 {
		if contains(array2, v) {
			alfa = append(alfa, v)
		}
	}
	for _, v := range array2 {
		if contains(array1, v) {
			alfa = append(alfa, v)
		}
	}
	fmt.Print("Наш алфавит \n")
	dublToZero(alfa)

	for i := 0; i < len(alfa); {
		if alfa[i] == 0 {
			alfa = append(alfa[:i], alfa[i+1:]...)
		} else {
			i++
		}
	}

	blyatNakhuyaEtoDelatSukaSuka(alfa)

}
func contains(s [12]rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func blyatNakhuyaEtoDelatSukaSuka(massiv []rune) {
	var massivString []string
	lenBlyat := len(massiv)
	var v string
	for i := 0; i < lenBlyat; i++ {
		v = string(massiv[i])
		massivString = append(massivString, v)
	}
	fmt.Println(strings.Join(massivString, ","))
}
func blyatNakhuyaEtoDelatSukaSuka12(massiv [12]rune) {

	var massivString []string
	lenBlyat := len(massiv)
	var v string
	for i := 0; i < lenBlyat; i++ {
		v = string(massiv[i])
		massivString = append(massivString, v)
	}
	fmt.Println(strings.Join(massivString, ","))
}
func dublToZero(xs []rune) {
	for i := 0; i < len(xs)-1; i++ {
		x := xs[i]
		if x == 0 {
			continue
		}
		for j := i + 1; j < len(xs); j++ {
			if xs[j] == x {
				xs[i] = 0
			}
		}
	}
}
