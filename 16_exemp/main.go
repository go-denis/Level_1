package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	var arr = []int{324, 234, 5, 6643, 345, 23, 29, 99, 67, 45}

	fmt.Println(quickSort(arr, 0, len(arr)-1))
}

func quickSort(arr []int, low, high int) []int {
	//Проверяем каждый элемент массива
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		//Уменьшаем значене стержня на 1, вызвываем рекрсию
		arr = quickSort(arr, low, p-1)
		//Увеличиваем минимальную точку на 1, вызвываем рекрсию
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

/*
Функция имеет один цикл for, который варьируется от самого низкого индекса до самого высокого индекса в массиве.
Сама по себе partition()функция есть O(n). Общая сложность быстрой сортировки зависит от того, сколько раз partition() вызывается.
*/
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] //Местоположене начальное стержня сортировки
	i := low           //Начальная точка сортировки
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	//Меняем местами
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
