package main

import (
  "fmt"
)

func compositeTypes() {
  nums := [3]int{1, 2, 3}
  fmt.Println(nums[1])

  words := []string{`ready`, `stedy`}
  words = append(words, `go`)
  fmt.Println(words[1])

  dict := map[string]int{"one": 1, "two": 2}
  fmt.Println(dict["one"])
}

type person struct {
  name string
  age int
}

func structs() {
  p := person{name: "Misha", age: 30}
  fmt.Println(p.name, p.age)
}

func main() {
  compositeTypes()
  structs()
}
