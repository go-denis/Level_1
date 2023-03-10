package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
var (
	mass   = []int{2, 4, 6, 8, 10} //Наш массив
	square int
	lock   sync.Mutex     //Для синхронной рабты горутин
	wg     sync.WaitGroup //Группа ожидания ожидает завершения набора горутинов.
)

func main() {

	/*
		Группа ожидания ожидает завершения набора горутинов. Основная подпрограмма вызывает Add,
		чтобы задать количество ожидаемых подпрограмм. Затем каждая из подпрограмм запускается и вызывает Done, когда закончит.
		В то же время ожидание можно использовать для блокировки до тех пор, пока не завершатся все горутины.
		Группа ожидания не должна копироваться после первого использования.
	*/

	wg.Add(len(mass)) //Сколько раз надо выполнить действие. Add, чтобы задать количество ожидаемых подпрограмм.
	for _, num := range mass {

		go Square(num) //Горутина

		fmt.Println("Основной поток")
	}
	wg.Wait() //Ждем завершения всех горутин

}

func Square(num int) {

	lock.Lock()          //Блокируем
	defer lock.Unlock()  //Открываем после выполнения
	*&square = num * num //Возведение в квадрат
	//Выводим в stdout Вывод будет в производном порядке, так как мы не знаем в какой момент запуситься горутина
	//При помощи каналов можно решить данную проблему, передавая от одной горутине в другую и создав условие.
	//Но тогда нагрузка ляжет на память, так как придется ждать ответа от памяти
	// и скорее всего будет лучше не использовать многопоточность, также действие
	//на возведение в степень выполняет не одно а где-то три действия, если смотреть низкоуровнево
	fmt.Println(square)
	wg.Done() //Декримент для счетчика Wait Group
}
