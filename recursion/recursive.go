package main

import "fmt"

func main() {
	fmt.Println(factorial(4))          // 24
	fmt.Println(fibonacci(1000000000)) // runtime: goroutine stack exceeds 1000000000-byte limit fatal error: stack overflow
}

// 求解 N 的阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2) // 存在大量重复计算，可用向后迭代优化，或矩阵优化
}

// 最为高效的尾递归
func fib(n int) int {
	return tailFib(n, 1, 1)
}

func tailFib(n, curSum, finalSum int) int {
	if n <= 2 {
		return finalSum
	}
	return tailFib(n-1, finalSum, curSum+finalSum) // 无计算
}
