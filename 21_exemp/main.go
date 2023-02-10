package main

import (
	"fmt"
)

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Интерфейс телефонов
// К примеру юзер хочет вставить в телефон андроид зарядку от айфона
type phone interface {
	insertPowerIphoneInAdnroid()
}

// У нас есть структура айфона,
type iphone struct{}

// Реализуем структуру методом
func (ip *iphone) insertPowerIphoneInAdnroid() {
	fmt.Println("Зарядка айфона началась")
}

// Структура адаптера
type AndroidApapter struct {
	androidPhone *android //Адаптируем под андроид
}

// Реализовываем метод адаптера, в который втыкаем шнур от айфона, и подключаем его к андроиду
func (a *AndroidApapter) insertPowerIphoneInAdnroid() {
	a.androidPhone.insertPowerAndroid() //Вставляем шнур андроид
}

// Структура андроид
type android struct{}

func (w *android) insertPowerAndroid() {
	fmt.Println("Зарядка андроид началась")
}

type user struct{}

// Реализуем структуру юзер через адаптер
func (u *user) InsertСhargingAndroidInIphone(com phone) {
	com.insertPowerIphoneInAdnroid()
}
func main() {
	user := &user{}     //Пользователь
	iphone := &iphone{} //Шнур айфон

	user.InsertСhargingAndroidInIphone(iphone) //Вставляем шнур от айфона

	android := &android{}

	androidAdapter := &AndroidApapter{
		androidPhone: android,
	}

	user.InsertСhargingAndroidInIphone(androidAdapter) //Вставляем шнур от айфона в андроид через адаптер
}
