package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
//Способ 1. С помощью встраивания

type Human struct {
	add  string
	name string
}

type Action struct {
	Human //Втраиваем
	get   string
}

//Метод структуры Human
func (h *Human) Add(add string) string {
	h.add = add
	return h.add
}
func main() {
	human := Human{add: "Отдал деньги", name: "Miki"}
	action := &Action{ //Записываем в переменную нашу структуру
		Human: human,
		get:   "Взял торт",
	}
	//action.Human.Add("Отдал монеты")
	fmt.Printf("%s, %s, %s", human.name, action.Human.Add("Добавил денег"), action.get)
}

//Способ 2. С помощью интерфейсов, так как Го не поддерживет наследование типов
/*
//Создаем наш интерфейс для связи структур с методом Add(), он содержит этот метод
type HumanAction interface {
	Add()
}

//Таже самая структура Human
type Human struct {
	add  string
	name string
}

//Структура Acnion
type Action struct {
	Human //Втраиваем
	get   string
}

//Метод принимает в качестве аргумента интерфейс
func chek(ha HumanAction) {
	ha.Add()
}

//Метод структуры Human
func (h *Human) Add() {

	fmt.Println(h.add) //Выводим действие на экран
}

func main() {
	human := Human{add: "Отдал деньги", name: "Miki"}
	action := &Action{ //Записываем в переменную нашу структуру
		Human: human,
		get:   "Взял торт",
	}
	//action.Human.Add("Отдал монеты")
	action.Add() //Косвено реализует метод от родительской структуры
	fmt.Println(action.name, action.get)
	chek(action) //Так как дочерняя структура косвенно реализует метод add, мы можем ее передать в функцию chek
}
*/
