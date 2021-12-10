package main

import (
	"fmt"

	"github.com/Direct-Dev-Ru/golang-module/calc"
)

func main() {
	result := calc.Add(1, 2)
	fmt.Println("calc.Add(1,2)==>", result)
}
