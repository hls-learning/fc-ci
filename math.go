package main

import "fmt"

func main() {
	fmt.Println(Soma(10,10))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtrai(a int, b int) int {
	return a - b
}

func Multiplica(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
