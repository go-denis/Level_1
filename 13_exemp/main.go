package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	var (
		x = 5
		y = 10
	)

	//Способ 1 с использованием арифметических операторов
	/*
		Идея состоит в том, чтобы получить сумму в одном из двух заданных чисел.
		Затем числа можно поменять местами, используя сумму и вычитание из суммы.
	*/
	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("\nСпособ 1: x = %d, y = %d", x, y)

	//Способ 2 Умножение и деление также можно использовать для замены.
	var (
		e = 5
		t = 10
	)
	e = e * t
	t = e / t
	e = e / t

	fmt.Printf("\nСпособ 2: x = %d, y = %d", e, t)

	//Способ 3
	/*
		Побитовый оператор XOR можно использовать для замены двух переменных.
		XOR двух чисел x и y возвращает число, все биты которого равны 1 везде, где биты x и y отличаются.
		Например, XOR из 10 (в двоичном 1010) и 5 (в двоичном 0101) равно 1111, а XOR из 7 (0111) и 5 (0101) равно (0010).
	*/
	var (
		q = 5
		w = 10
	)
	q = q ^ w // x now becomes 15 (1111)
	w = q ^ w // y becomes 10 (1010)
	q = q ^ w // x becomes 5 (0101)
	fmt.Printf("\nСпособ 3: x = %d, y = %d", q, w)
}
