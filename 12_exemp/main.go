package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
func main() {
	var (
		str  = []string{"cat", "cat", "dog", "cat", "tree"}
		mapa = make(map[string]struct{})
	)

	//Записываем в множество
	for _, v := range str {
		mapa[v] = struct{}{}
	}

	fmt.Println(mapa)

}
