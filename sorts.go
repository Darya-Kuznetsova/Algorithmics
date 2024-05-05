package main

import "fmt"

func main() {
	array_1 := [5]int{6, 9, 4, 0, 8}
	bubble_sort(array_1[:])
	array_2 := [6]int{7, 8, 5, 4, 0, 0}
	bubble_sort_version_2(array_2[:])
	array_3 := [6]int{5, 8, 9, 3, 2, 0}
	direct_sort(array_3[:])
	array_4 := [6]int{7, 0, 0, 3, 12, 1}
	insert_sort(array_4[:])
	array_5 := [7]int{9, 5, 0, 1, 2, 1, 12}
	fmt.Println(quick_sort(array_5[:], 0, len(array_5)-1))
	array_6 := [6]int{6, 33, 8, 0, 3, 17}
	fmt.Println(heap_sort(array_6[:]))

}

/*
BUBBLE SORT: "самые легкие всплывают первыми".
Сравниваем элемент с рядом стоящим:
array [0] = 9:
1. [9, 5, 7, 4, 3, 2, 1, 6, 8, 0] -> [5, 7, 4, 3, 2, 1, 6, 8, 0, 9]
2. [5, 4, 3, 2, 1, 6, 7]
3. [5, 4, 3, 2, 1, 6, 7, 0, 8, 9] etc.
O(n ^ 2), не самая производительная сортировка
*/

func bubble_sort(array []int) {
	for i := 0; i <= len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			/* len(array) - 1 т.к. array[len(array) - 2] сравнивается с array[len(array) - 1],
			а array[len(array) - 1] уже ни с чем не сравнивается */
			if array[j] > array[j+1] {
				memory := array[j]
				array[j] = array[j+1]
				array[j+1] = memory
			}
		}
	}
	fmt.Println(array)

}

/* for - if!(condition) break - замена do-while
do-while: condition выполняется -> цикл продолжается
for-if!(condition) break: !condition выполняется -> цикл break
*/
func bubble_sort_version_2(array []int) {
	var finish bool
	for {
		finish = true
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				memory := array[i]
				array[i] = array[i+1]
				array[i+1] = memory
				finish = false
			}
		}
		if finish {
			break
		}
	}
	fmt.Println(array)

}

/*
DIRECT SORT (сортировка выбором): начинаем с нулевого элемента и сравниваем его со всеми остальными, находя самый меньший.
Так же O(n^2), как и пузырьковая
*/

func direct_sort(array []int) {
	// Проход по всему массиву:
	for i := 0; i < len(array)-1; i++ {
		min_position := i
		// Все остальные элементы массива с которыми сравнивается i:
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_position] {
				min_position = j
			}
		}
		if i != min_position {
			memory := array[i]
			array[i] = array[min_position]
			array[min_position] = memory
		}
	}
	fmt.Println(array)
}

/*
INSER SORT (сортировка вставками): элемент, начиная с нулевого сравнивается с каждым последующим
и каждый раз обменивается
And again O(n^2)
*/
func insert_sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				memory := array[i]
				array[i] = array[j]
				array[j] = memory
			}
		}
	}
	fmt.Println(array)
}

/*
QUCK SORT - продвинутая сортировка. Разбивка на два подмассива, центральная точка - pivot.
Разбивка происходит до тех пор, пока все не будет отсортировано.
O(log n) EXEPT: if pivot - крайний элемент в уже отсортированном массиве, тогда O(n)
Обычно, этот алгоритм использует стандартная библиотека большинства ЯП
*/

func quick_sort(array []int, start_position, end_position int) []int {
	left_position := start_position
	right_position := end_position
	pivot := array[(start_position+end_position)/2] // another математическая формула нахождения середины
	for left_position <= right_position {
		for array[left_position] < pivot {
			left_position++
		}
		for array[right_position] > pivot {
			right_position--
		}
		if left_position <= right_position {
			if left_position < right_position {
				memory := array[left_position]
				array[left_position] = array[right_position]
				array[right_position] = memory
			}
			left_position++
			right_position--
		}

	}
	if left_position < end_position {
		quick_sort(array, left_position, end_position)
	}
	if start_position < right_position {
		quick_sort(array, start_position, right_position)
	}
	return array

}

/*
BINARY HEAP (бинарная куча) - у каждого элемента двое детей (и один родитель соответственно).
Элемент у которого нет родителя - корень.
Если родитель имеет индекс i, то дочерние 2 * i + 1 and 2 * i + 2
HEAP SORT (пирамидальная сортировка или кучей) - сравнение родительского элемента с дочерними.
Например, для сортировки от мин к макс используется восходящая куча, когда родитель больше детей.
Используется реже и медленненее, чем quick (для решения ситуацииб когда quick может дать O(n))
*/

func heap_sort(array []int) []int {
	// подготовка кучи (поиск максимального элемента):
	for i := len(array)/2 - 1; i >= 0; i-- { // len(array) / 2 - 1 середина
		// просеивание(самый большой элемент будет на нулевом индексе)
		heapify(array, len(array), i)
	}
	// элементы из кучи:
	for i := len(array) - 1; i >= 0; i-- {
		// перемещение текущего корня (самого большого) в конец:
		memory := array[0]
		array[0] = array[i]
		array[i] = memory

		// просеивание для уменьшенной куче без последнего элемента:
		heapify(array, i, 0)
	}
	return array
}

//Операция просеивания - выносит наверх элементы, которые не являются самыми большими:

func heapify(array []int, heap_size, root_index int) {
	largest := root_index // наибольший элемент
	// определяем дочерние элементы:
	left_child := 2*root_index + 1
	right_child := 2*root_index + 2
	if left_child < heap_size && array[left_child] > array[largest] {
		largest = left_child
	}
	if right_child < heap_size && array[right_child] > array[largest] {
		largest = right_child
	}
	if largest != root_index {
		memory := array[root_index]
		array[root_index] = array[largest]
		array[largest] = memory

		// recursion для других нижележащих элементов:
		heapify(array, heap_size, largest) // меняем root на largest
	}

}
