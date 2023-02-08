package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// Создаем потокобезопасную map, с помошью встраивания в мьютекс
type SafeMap struct {
	mx   sync.Mutex
	mapa map[int]int
}

// Конструктор для мапы
func NewMapa() *SafeMap {
	return &SafeMap{
		mapa: make(map[int]int),
	}
}

// Функция - обертка для записи
func (sm *SafeMap) Store(key int, value int) {
	sm.mx.Lock()
	/*
		defer имеет небольшой оверхед (порядка 50-100 наносекунд), поэтому если у вас код
		для высоконагруженной системы и 100 наносекунд имеют значение, то вам может быть выгодней не использовать defer
	*/
	defer sm.mx.Unlock()
	sm.mapa[key] = value
}

func (sm *SafeMap) PrintMap() {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	for key, val := range sm.mapa {
		fmt.Printf("\nКлюч - %d, Значение - %d", key, val)
	}
}

func main() {

	mapa := NewMapa()

	go func() {
		for {
			/*
				Молимся чтобы рандомные значения ключей в рутинх не свопали
			*/
			key := rand.Int()   //Генерируем ключ
			value := rand.Int() //Генерируем значение
			mapa.Store(key, value)

		}

	}()

	go func() {
		for {
			/*
				Молимся чтобы рандомные значения ключей в рутинх не свопали
			*/
			key := rand.Int()   //Генерируем ключ
			value := rand.Int() //Генерируем значение
			mapa.Store(key, value)
		}

	}()

	fmt.Scanln()
	mapa.PrintMap()
}
