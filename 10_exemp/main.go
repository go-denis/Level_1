package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc
*/
var (
	mapa = make(map[int]map[float64]struct{})
	mass = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0}
)

func main() {

	for _, v := range mass {

		num := 0
		if v > 0 {

			num = int(math.Floor(v/10) * 10)

		} else if v == 0 {

			num = 10
		} else {

			num = int(v/10) * 10

		}

		//Если нет такого ключа
		if _, ok := mapa[num]; !ok {

			mapa[num] = make(map[float64]struct{})
		}
		//Добавляем значение в карту
		mapa[num][v] = struct{}{}

	}

	for key, v := range mapa {
		s := "" + fmt.Sprintf("%d", key) + " :{"
		//fmt.Printf("%d :{", key)
		i := 0
		for g := range v {
			if i > 0 {
				s += ", "
			}

			s += "" + fmt.Sprintf("%.1f", g) + ", "
		}
		s = s[:len(s)-2] //Убераем последнюю запятую
		s += "} "

		fmt.Println(s)
	}

}

// Функция для округления к меньшему числу
/*
func round(x float64) float64 {
	if math.IsNaN(x) {
		return x
	}
	if x == 0.0 {
		return x
	}
	roundFn := math.Ceil
	if math.Signbit(x) {
		roundFn = math.Floor
	}
	xOrig := x
	x -= math.Copysign(0.5, x)
	if x == 0 || math.Signbit(x) != math.Signbit(xOrig) {
		return math.Copysign(0.0, xOrig)
	}
	if x == xOrig-math.Copysign(1.0, x) {
		return xOrig
	}
	r := roundFn(x)
	if r != x {
		return r
	}
	return roundFn(x*0.5) * 2.0
}
*/
