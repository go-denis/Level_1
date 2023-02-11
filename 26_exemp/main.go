package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

*/

func main() {

	str := "AdcfDA" //"asdfghjk" "AdcfDA"

	OriginalCheck(str)
}

func OriginalCheck(str string) {
	var (
		strCheckMap      = make(map[int]string) //Пытался со слайсом, но это геморройно немного
		check       bool = true                 //Для проверки
	)

	str = strings.ToLower(str) //Переводим в нижний регистр
	strRun := []rune(str)      //Переводим в руны для перебора

	//Перебирваем руны
	for i, v := range strRun {
		v := string(v) //Форматируем в строку

		for _, sv := range strCheckMap {
			if v == sv {
				fmt.Printf("\n%s, %s", v, sv)
				*&check = false
			}
		}
		strCheckMap[i] = string(v) //Записываем в новый славс рун

	}
	if *&check == false {
		fmt.Printf("'%s', строка не уникальна, есть повторения %t", *&str, *&check)
	} else {
		fmt.Printf("'%s', строка уникальна", *&str)
	}
}
