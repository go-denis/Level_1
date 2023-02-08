package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

*/

func main() {
	var (
		ch1  = make(chan int)
		ch2  = make(chan int)
		wg   sync.WaitGroup
		mass = [10]int{5, 6, 4, 5, 1, 66, 5, 45, 23, 98}
	)

	wg.Add(3)
	//Пишем значени в первый канал
	go func() {
		for _, v := range mass {
			ch1 <- v
		}
		close(ch1)
		wg.Done()
	}()

	//Пишем в первый канал произведение чисел из первого канала
	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}
		close(ch2)
		wg.Done()
	}()

	//Читаем и выводим в консоль данные из второго канала
	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Wait()

}
