package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	mx sync.Mutex //Для защиты счетчика
	i  int        //Счетчик
}

func main() {
	var (
		wg      sync.WaitGroup
		counter = Counter{
			i: 0,
		}
	)

	wg.Add(5) //Запускаем 10 рутин

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 5; i++ {
				counter.mx.Lock() //Блокируем доступ к переменной

				counter.i++         //Инкрементируем счетсик
				counter.mx.Unlock() //Разбокируем доступ к переменной
			}
			wg.Done() //Рутина корректно отработала
		}()

	}
	//Ждем завершение всех рутин
	wg.Wait()

	fmt.Println(counter.i)
}
