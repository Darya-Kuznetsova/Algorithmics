package main

import "fmt"

// При оценке сложности алгоритмов используется нотация O(f(n)), n - размер входных данных.
// Не реальное количество шагов, а для примерной оценки.
// Для линейной зависимости O(n), для квадратичной зависимости O(n ^ 2)
// O(1) константная сложность (всегда за одно время) - например, поиск элемента в массиве по индексу
// Обычно (бывает обычный алгоритм быстрее) лучше выбирать константные алгоритмы (их мало) of course
// если есть возможность.

func main() {
	fmt.Println(find_available_divider(12)) // [1 2 3 4 6 12]
	fmt.Println(simple_numbers(12))         // [1 2 3 5 7 11]
	fmt.Println(sum(7))                     // 28
}

// 1. Находит все допустимые делители числа:

func find_available_divider(number int) []int {
	result := make([]int, 0)
	// Кол-во делений числа зависит от этого числа:
	// Данная зависимость является ЛИНЕЙНОЙ
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

// 2. Простые числа [1, n] (делятся на 1 и на себя):

func simple_numbers(n int) []int {
	result := make([]int, 0)
	// КВАДРАТИЧНАЯ зависимость:
	// [1, n]
	for i := 1; i <= n; i++ {
		simple := true
		// Вcе делители: [2, n]
		for j := 2; j < i; j++ {
			if i%j == 0 {
				simple = false
			}
		}
		if simple {
			result = append(result, i)
		}
	}
	return result
}

// Cумма чисел от 1 до n:

func sum(number int) int {
	result := 0
	for i := 1; i <= number; i++ {
		result += i
	}
	return result
}
