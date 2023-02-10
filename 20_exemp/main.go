package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	var str = "snow dog sun"
	fmt.Println(reverseWords(str))
}
func reverseWords(s string) string {
	//Разбиваем строку по пробелам в в массив строк
	words := strings.Fields(s)
	fmt.Println(words)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	//Join объединяет элементы своего первого аргумента для создания одной строки.
	return strings.Join(words, " ")
}
