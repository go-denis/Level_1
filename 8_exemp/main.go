package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
*/

// подсмотрел решение так как не знал как делать, но разобрался
func main() {
	var (
		num  = rand.Int63()
		nBit = 4
	)
	//Устанавливаем 1
	numAdd := AddOne(num, nBit)
	numZero := AddZero(num, nBit)
	//Используем %b для вывода именно битового значения
	fmt.Printf("Было \t %b", *(*uint64)(unsafe.Pointer(&num))) //переводим в двоичный формат
	fmt.Printf("\nУстанавливаем 1\t %b", *(*uint64)(unsafe.Pointer(&numAdd)))
	fmt.Printf("\nУстанавливаем 0\t %b", *(*uint64)(unsafe.Pointer(&numZero)))
}

func AddOne(num int64, pointer int) int64 {
	numAdd := num | 1<<(pointer-1) //Отнимаем 1 чтобы вернуть в нужное положение
	return numAdd                  //num | 1<<(pointer-1)
}

func AddZero(num int64, pointer int) int64 {
	//numAdd := num | 1<<pointer
	return num &^ (1 << (pointer - 1))
}
