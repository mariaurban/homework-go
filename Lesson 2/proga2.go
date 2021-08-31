package main

import (
  "fmt"
)

func main()  {
  fmt.Println("Здарова, человек. Введи свое имя")
  var name string
  fmt.Scanln(&name)
  fmt.Println("Привет", name,"! Удачного тебе дня :)")
}
