package main

//5. Задача с 3 картами
import (
	"fmt"
	"math/rand"
)

func main() {
	percent := 0
	fmt.Println("Человек, введи  количество циклов")
	var numberOfCycles int
	fmt.Scanln(&numberOfCycles)
	for round := 0; round < numberOfCycles; round++ {
		i := 1 + rand.Intn(6)
		if i == 3 || i == 4 {
			percent++
		}
	}
	percent1 := float64(percent)
	numberOfCycles1 := float64(numberOfCycles)
	fmt.Println("Если вы будете играть в эту странную игру, то ваш провент выигрыша будет равен :", (percent1/numberOfCycles1)*100, "%")

}
