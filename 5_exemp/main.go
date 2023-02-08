package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	var (
		ch       = make(chan int)
		stopTime = 5
		mutex    sync.RWMutex //Определяем мьютекс для блокирровки только чтения
	)

	//Отправляет значения в канал
	go func() {

		for {
			mutex.RLock()                    //Блокируем чтение
			defer mutex.RUnlock()            //Разблокируем
			time.Sleep(1 * time.Millisecond) //Заставляем ждать, чтобы отправка выполнялась последовательно
			data := rand.Int()               //Генерируем случайные числа
			ch <- data                       //Записываем данные
			fmt.Println("\nЗаписал -", data) //Читаем

		}
	}()

	//Читает значения из канала
	go func() {

		for data := range ch {
			mutex.RLock()                   //Блокируем чтение
			defer mutex.RUnlock()           //Разблокируем
			fmt.Println("\nСчитал -", data) //Выводим в термнал

		}

	}()
	//Заставляем ждать главный поток
	time.Sleep(time.Duration(stopTime) * time.Second)

}
