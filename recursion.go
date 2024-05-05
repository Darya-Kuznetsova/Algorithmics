package main

import "fmt"

// RECURSION: пример ЭКСПОНЕНТАЛЬНОЙ зависимости O(2 ^ n) - при увеличении на 1 сложность увеличивается на 2:

func main() {
	fmt.Println(fibonacci(7))     // 13
	fmt.Println(fibonacciLine(8)) // 13
	fmt.Println(factorial(5))     // 120

}

// Fibonacci RECURSION

func fibonacci(index int) int {
	if index == 1 || index == 2 {
		return 1
	}
	return fibonacci(index-1) + fibonacci(index-2)

}

// FIbonacci LOOP:

func fibonacciLine(index int) int {
	if index == 1 || index == 2 {
		return 1
	} else {
		slice := make([]int, index)
		slice[0] = 0
		slice[1] = 1
		for i := 2; i < index; i++ {
			slice[i] = slice[i-1] + slice[i-2]
		}
		return slice[index-1]

	}
}

// Factorial RECURSION:

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}
