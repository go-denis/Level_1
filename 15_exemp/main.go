package main

import "fmt"

var justString string

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/
/*
1 Стоит добавить метод createHugeString
 Слайс байт строки v останется в памяти, но не будет доступен,
 т.к. v недоступка вне someFunc
*/
func createHugeString(num int) string {
	fmt.Println(num)
	return "ksdjdieывааыаыаыаыаыаыаываывучцуууцагукрагшаршаоцщаощзуаозцщоащзцуаошщзаоцщзаощзаоцущзаоцузщацузщаощцуаощзцуоащцуощщощахзмдывмбььд"
}

func someFunc() {
	v := createHugeString(1 << 10)
	//fmt.Println(v)

	//Задесь мы видимо хотим взять подстроку, длинной 100 символов, но к примеру если у нас меньше то будет ошибка
	//Но здесть мы получим строку из слайса байт
	//Если уменьшить на 1
	//При этом я сделал строку более 100 символов
	//justString = v[:99] //то все ок
	//fmt.Println(justString)

	//Нужна проверка на кол-во элементов в слайсе
	rjs := make([]rune, 100)
	i := 0
	for _, r := range v {
		if i >= 100 {
			break
		}
		rjs[i] = r
		i++
	}
	justString = string(rjs)
	fmt.Println(justString)

	// Взять слайс первых 4 байт
	fmt.Println(v[:100])

	// Взять слайс первых 5 байт
	fmt.Println(v[:99])
}

func main() {
	someFunc()
}
