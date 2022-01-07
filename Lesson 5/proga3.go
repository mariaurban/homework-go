package main

import "fmt"

// Алфавит: Есть два массива символов.
// Собрать третий массив из символов которые есть и в том и в другом массиве.
// Каждый символ в результате не упоминается больше одного раза
// Вывести третий массив

func main() {
	array1 := [12]string{"2", "4", "2", "f", "q", "e", "2", "s", "D", "w", "6", "3"}
	array2 := [12]string{"3", "2", "7", "6", "d", "a", "w", "e", "w", "?", "p", "l"}
	var alfa []string
	for _, v := range array1 {
		if !contains(alfa, v) {
			alfa = append(alfa, v)
		}
	}
	for _, v := range array2 {
		if !contains(alfa, v) {
			alfa = append(alfa, v)
		}
	}
	fmt.Print("Наш алфавит \n")
	blyatNakhuyaEtoDelatSukaSuka(alfa)

}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func blyatNakhuyaEtoDelatSukaSuka(massiv []string) {
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		fmt.Println(massiv[i])
	}
}
