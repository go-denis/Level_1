package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}
*/

func main() {
	user := []interface{}{5, "asda", false, make(chan int)}

	DetectUserType(user)
}

func DetectUserType(user []interface{}) {

	//Определяем тип
	for _, w := range user {
		fmt.Printf("\nPrintf %%T:\t%T\n", w)
		fmt.Println("reflect.TypeOf \t", reflect.TypeOf(w))
		fmt.Println("reflect.ValueOf.Kind:\t", reflect.ValueOf(w).Kind())

	}
}
