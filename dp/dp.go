package main

import "fmt"

func main() {
	fmt.Println(fib(3), fib(4))   // 3 5
	fmt.Println(climbStairs1(10)) // 89
	fmt.Println(climbStairs2(10)) // 89
}

// 记忆化搜索
func fib(n int) int {
	m := make([]int, n+1)
	m[0], m[1] = 1, 1
	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

// 借助数组保存中间结果，避免重复计算
// 时间 O(N)，空间 O(N)
func climbStairs1(n int) int {
	steps := make([]int, n+1)
	steps[1], steps[2] = 1, 2 // 边界

	for i := 3; i <= n; i++ {
		steps[i] = steps[i-2] + steps[i-1] // 状态转移
	}

	return steps[n] // 最优子结构
}

// 向后累计步数
// 时间 O(N)，空间 O(1)
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}

	i, j := 1, 2
	for n > 2 {
		i, j = j, i+j // bottom -> up
		n--
	}
	return j
}
