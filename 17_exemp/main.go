package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
/*
Двоичный (бинарный) поиск (также известен как метод деления пополам или дихотомия) —
классический алгоритм поиска элемента в отсортированном массиве (векторе), использующий дробление массива на половины
*/
func main() {
	var arr = []int{324, 234, 5, 6643, 345, 23, 29, 99, 67, 45}

	fmt.Println(binarySearch(arr, 5))
}
func binarySearch(a []int, search int) (result int, searchCount int) {
	//Делем массив на два
	mid := len(a) / 2

	switch {
	case len(a) == 0:
		result = -1 //Не нашел
	case a[mid] > search:
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // Если что-то, кроме -1 не найдено
			result += mid + 1
		}
	default:
		result = mid // Нашел!
	}

	searchCount++
	return
}
