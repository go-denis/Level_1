package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры
Point с инкапсулированными параметрами x,y и конструктором.
*/

// Немного не допонял что означает инкапусированные параметры и конструктор в данной задаче
// Поэтому сделаю рассчет по формуле для
type Point struct {
	x float64
	y float64
}

// Возможно про этот конструктор идет речь
func ConfigPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Собираем конструктором Х
func (p *Point) X() float64 {
	return p.x
}

// Собираем конструктором Y
func (p *Point) Y() float64 {
	return p.y
}

// d = √ (х А – х В) 2 + (у А – у В) 2. Формула расчета
func distance(x, y Point) float64 {
	katet1 := math.Sqrt(math.Pow(math.Abs(x.X()-y.Y()), 2)) //Находим первый катет
	katet2 := math.Sqrt(math.Pow(math.Abs(x.Y()-y.Y()), 2)) //Находим второй катет
	d := math.Ceil(katet1 + katet2)
	return d
}

func main() {
	point1 := Point{x: 55.2, y: 77.33}
	point2 := Point{x: 61.666, y: 99.23}

	fmt.Println("Расстояние между x и y на плоскости = ", distance(point1, point2))
}
