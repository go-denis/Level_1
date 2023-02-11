package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	Sleep(30.3)
	fmt.Println("Ok")
}
func Sleep(sec float64) {
	t := time.Now()
	/*if time.Now().Sub(t).Seconds() < sec {
		Sleep(sec)
	}*/
	for time.Now().Sub(t).Seconds() < float64(sec) {
	}
	fmt.Println("Время вышло!")

}
