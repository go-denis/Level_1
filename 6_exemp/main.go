package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	var (
		ch          = make(chan int)
		stop        = make(chan bool, 1)                       //С помощью буферизированного канала, также с помощью не буферезировнного, таже будет выглядеть
		ctx, cancel = context.WithCancel(context.Background()) //С помощью контекста
		wg          sync.WaitGroup                             //Спомощью группы ожидания
	)

	//Горутина, завершает работу по записи в доп канал
	wg.Add(1)
	go func(ctx context.Context) {

		for {
			select {
			case <-stop:
				fmt.Println("Остановка каналом")
				wg.Done()
				return
			default:

				fmt.Println("Я работаю еще")
			}
			time.Sleep(500 * time.Millisecond)

		}

	}(ctx)
	time.Sleep(500 * time.Microsecond)
	stop <- true //Останавливаем каналом
	wg.Wait()

	//Горутина завершающая свою работу при помощи контекста
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка контекстом")
				wg.Done()
				return
			default:

				fmt.Println("И я работаю")
			}
		}

	}(ctx)

	time.Sleep(500 * time.Microsecond)
	cancel()  //Оставнавливаем контекстом
	wg.Wait() //Ожидаем завершения

	//Горутина завершает работу по условию
	wg.Add(1)
	ok := true
	go func() {
		for {

			if ok != false {
				fmt.Println("Остановка по условию")
				wg.Done() //Информируем о завершении
				return
			}
			fmt.Println("Так я тоже работаю")
		}
	}()
	time.Sleep(500 * time.Microsecond)
	ok = false
	wg.Wait() //Ожидаем завершения

	/*
		Горутина Завершит чтение, если канал остановят

		Закрытием канала можно завершить сразу несколько горутин
	*/
	for i := 0; i < 300; i++ {
		ch <- 123 //Записываем в канал
	}
	wg.Add(1)
	go func() {
		for read := range ch {
			fmt.Println("А я читаю - ", read)
		}
		wg.Done() //Информируем о завершении
	}()

	//Немного подождем
	time.Sleep(100000 * time.Microsecond)
	close(ch) //Закрываем канал
	wg.Wait() //Ожидаем завершения

	fmt.Scanln()
	//time.Sleep(10000 * time.Microsecond)
}
