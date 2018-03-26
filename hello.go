package main

import (
  "fmt"
)

func main() {
  test := "Hello"
  phrase1, phrase2 := doSomething(test)
  
  fmt.Println(phrase1)
  fmt.Println(phrase2)
}

func doSomething(x string) (string, string)  {
  return (x + " Jim"), ("Rock On")
}