package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	//Генерируем множества
	map1 := GenSets(22, 44)
	map2 := GenSets(22, 22)
	//Выводим их на экран
	fmt.Println("map1:", StringConvert(map1))
	fmt.Println("map2:", StringConvert(map2))
	//Создаем новое множество  с пересечениями множеств
	map3 := Intersect(map1, map2)
	fmt.Println("intersect:", StringConvert(map3))
}

// GenSets добавляет в возвращаемое множество n in [0,k) случайных чисел
func GenSets(k int, n int) *map[int]struct{} {

	set := make(map[int]struct{})
	for i := 0; i < n; i++ {
		//Добавление в карту
		set[rand.Intn(k)] = struct{}{}
	}
	return &set
}

// Записываем множество в строку, чтобы вывести на экран
func StringConvert(set *map[int]struct{}) string {
	var s []string
	for v := range *set {
		s = append(s, strconv.Itoa(v))
	}
	return "{" + strings.Join(s, " ") + "}"
}

// Вычисляем пересечение множеств и возвращаем его
func Intersect(map1 *map[int]struct{}, map2 *map[int]struct{}) *map[int]struct{} {
	set := make(map[int]struct{})

	// Проверяем с помощью мап меньшего размера
	if len(*map1) > len(*map2) {
		map1, map2 = map2, map1
	}

	// Перебираем мапу1 чтобы найти сходства в мапе2 и записать в мапу3 пересечения множеств
	for key := range *map1 {
		//Если в мапе2 содержиться число из мапы1, то записываем в новую мапу с пересечениями
		if _, ok := (*map2)[key]; ok {
			set[key] = struct{}{}
		}
	}

	return &set
}
