package main

import (
	"fmt"
	"math/rand"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

*/
//Как я понял из условия числа а и b должны быть больше 2 или не 20
func main() {
	var (
		a int = rand.Int() * 3 / -1
		b int = rand.Int() * 3 / -1
	)
	fmt.Println(a, b)
	//flag.Int(&a, "a", int(math.Pow(2, 20)+1), "a")
	//flag.Int(&b, "b", int(math.Pow(2, 20)+1), "b")
	//flag.Parse()

	fmt.Printf("%d \n %d \n %d \n %d", a*b, a/b, a+b, a-b)
}
