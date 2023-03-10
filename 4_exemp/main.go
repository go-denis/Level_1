package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	var (
		works      int = 10                      //Колличество воркеров
		jobs           = make(chan int)          //Канал для работы
		signalChan     = make(chan os.Signal, 1) //Буферизированный кнал для отслеживания завершения работы
		//idleConnsClosed     = make(chan bool)         //Флаг хорошего завершения работы
	)

	//Отслеживаем нажатие Ctrl+C
	signal.Notify(signalChan, os.Interrupt)

	go Worker(works, jobs)

	//постоянныая запись в канал в главном потоке
	for {

		select {
		//Ждем когда канальная операция будет разблокированна
		case <-signalChan:
			/*
				При помощи закрытия канала мы обеспеечиваем немедленное прекращение работы всех воркеров
				А выходя из цикла, завершаем работу главного потока по бесконечной штамповке воркеров
			*/
			close(jobs) //Закрываем канал тем самым все горутины завершают свою работу
			return      //Выходим из цикла
			//пока не нажали Ctrl+C, продолжается запись
		default:
			jobs <- rand.Int()

		}

	}

}

func Worker(works int, job chan int) {

	for i := 0; i < works; i++ {
		fmt.Printf("\nВоркер №%d", i)
		//Генерируем горутины, столько раз сколько указанно в переменной works
		go func() {
			//Читаем из канала данные
			for j := range job {
				fmt.Printf("\n Cчитал - %d", j)

			}
		}()
	}

}
