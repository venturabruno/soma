package main

import (
	"fmt"
)

func soma(arg1 int, arg2 int) int{
	return arg1 + arg2
}

func main() {
	arg1 := 5
	arg2 := 5
	result := soma(arg1, arg2)
	fmt.Println("A soma de ", arg1, "+", arg2, "=", result)
}