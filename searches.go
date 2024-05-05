package main

import "fmt"

func main() {
	array_1 := [7]int{8, 12, 9, 12, 0, 4, 6}
	fmt.Println(simple_search(array_1[:], 12))
	array_2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(binary_search(array_2[:], 2, 0, len(array_2)-1))
}

/* Simple search: простой перебор элементов массива в поиске нужного*/

func simple_search(array []int, number int) []int {
	position := make([]int, 0)
	for i := 0; i < len(array); i++ {
		if array[i] == number {
			position = append(position, i)
		}
	}
	return position
}

/* Binary search - поиск в отсортированном массиве: */

func binary_search(array []int, number, min, max int) int {
	/* Определяем центральную точку относительно которой будет сранение
	Важно сделать проверку, так как при использовании мат. формулы (см. ниже)
	может быть перекрещивание max и min */

	var mid_point int
	if max < min {
		return -1
	} else {
		mid_point = (max-min)/2 + min // математическая формула поиска середины
	}
	if array[mid_point] < number {
		return binary_search(array, number, mid_point+1, max)
	}
	if array[mid_point] > number {
		return binary_search(array, number, min, mid_point-1)
	} else {
		return mid_point
	}
}
