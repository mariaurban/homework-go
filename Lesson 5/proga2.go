package main

import "fmt"

// Есть массив символов.
// Собрать из них строку, но в нее добавлять только символы латинского алфавита, пробелы и точки.
// Вывести результирующую строку.

func main() {
	someRandomSymbol := [11]rune{'y', 'w', 'W', ',', '.', '1', '3', '1', 'f', 'S', ' '}
	var newRandomSymbol []string
	var name string
	for i, _ := range someRandomSymbol {
		n := int(someRandomSymbol[i])
		name = string(n)
		if ((n < 90) && (n > 65)) || ((n < 122) && (n > 95)) || (n == 32) || (n == 46) {
			newRandomSymbol = append(newRandomSymbol, name)
		}

	}
	fmt.Print("Наш новенький массив выглядит так :\n")
	blyatNakhuyaEtoDelatSuka(newRandomSymbol)
}
func blyatNakhuyaEtoDelatSuka(massiv []string) {
	lenBlyat := len(massiv)
	for i := 0; i < lenBlyat; i++ {
		fmt.Println(massiv[i])
	}
}
