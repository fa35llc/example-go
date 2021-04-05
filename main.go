package main

import (
	"fmt"

	"example.com/example/calc"
)

type animal struct {
	kind    string
	habitat string
}

func (a *animal) setAnimal(k, h string) *animal {
	a.kind = k
	a.habitat = h
	return a
}

func main() {
	a := new(animal)
	r := a.setAnimal("猫", "犬")
	fmt.Println("動物", r)
	fmt.Println("r変数のアドレス", &r)

	s, err := calc.Sum(1, 4)
	if err != nil {
		fmt.Println("Error:Func sum:", err)
	}
	fmt.Println("Func sum:", s)

}
