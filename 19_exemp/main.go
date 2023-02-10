package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	var (
		str = "Denis Simonov"
	)
	newStr := Reverse(str)

	fmt.Println(newStr)
}
func Reverse(s string) string {
	runes := []rune(s) //Праедставляем строку в виде литерала рун
	//Перебираем руну
	for i, v := 0, len(runes)-1; i < v; i, v = i+1, v-1 {
		fmt.Println("i - ", i)
		fmt.Println("v - ", v)
		//меняем местами
		runes[i], runes[v] = runes[v], runes[i]
	}
	//Возвращаем с переводом в тип стринг
	return string(runes)
}
