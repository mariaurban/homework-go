package main

import (
  "fmt"
)

func main()  {
  fmt.Println("Человек, введи число")
  var num int
  fmt.Scanln(&num)
  fmt.Println("Ваше число", num)
}
