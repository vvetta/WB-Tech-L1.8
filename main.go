package main

import "fmt"

func main() {
	var number int64
	fmt.Println("Ведите переменную формата int64: ")	
	fmt.Scan(&number)

	var bitPos int
	fmt.Println("Введите бит, который хотите заменить: ")
	fmt.Scan(&bitPos)

	var bitVal int
	fmt.Println("Введите новое значение бита: ")
	fmt.Scan(&bitVal)

	if bitVal > 1 || bitVal < 0 { return }

	var newNumber int64
	newNumber = setBit(number, bitPos, bitVal)
	
	fmt.Printf("Новое число: %d", newNumber)
}

func setBit(number int64, bit int, bitVal int) int64 {
	if bitVal == 1 { 
		number |= (1 << bit)
	} else {
		number &^= (1 << bit)
	}

	return number
}
