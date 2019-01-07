### 关于递归

直观上看就是函数直接或间接地调用了自身。

递归求解复杂问题的 3 个过程：

- 将不好下手求解的大问题划分成性质相同的子问题
- 求解小的子问题
- 子问题的解回归组合成大问题的解

递归求解的 3 个关键：

- 求解递归式 `g(x)` 能将原问题分解为子问题：使 `f(x) = g(f(x-1))` 成立
- 递归必有终止条件：求解最小子问题时递归结束，可能有多个出口

- 递归必向终止条件逼近



### N 的阶乘

可直接迭代实现：

```go
func factorial(n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= n
	}
	return res
}
```

很容易发现，求解 N 的阶乘等于 **N * N-1的阶乘**（子问题），直到 N-1 到达终止值 0 其阶乘为 1。则有递归实现：

```go
func factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * factorial(n-1)
}
```

如求 `5!` 的递归过程：

<img src="https://images.yinzige.com/2018-12-27-090610.png" width="360px"/>

图片来源：[recursion-and-backtracking](https://www.hackerearth.com/zh/practice/basic-programming/recursion/recursion-and-backtracking/tutorial/)



### 关于尾递归

递归调用在计算机内部使用栈实现，每向下一层栈就加一帧内存，可能会出现栈溢出。如计算斐波那契数列：

```go
func main() {
	// runtime: goroutine stack exceeds 1000000000-byte limit fatal error: stack overflow
    fmt.Println(fibonacci(1000000000))
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2) // 存在大量重复计算，可用向后迭代优化，或矩阵优化
}
```

栈溢出是因为 `fibonacci(n-1)` 计算完后结果需暂时保存，等待 `fibonacci(n-2)` 计算完毕再求和，过多地保存子问题的解导致栈溢出，可用尾递归优化。

尾递归指函数能调用本身，但在 return 时不包含计算表达式，编译器能始终使用同一函数计算来求解子问题：

```go
func fib(n int) int {
	return tailFib(n, 1, 1)
}

func tailFib(n, curSum, finalSum int) int {
	if n <= 2 {
		return finalSum
	}
     // return 语句中没有需要保存的中间态，始终只占一层栈帧
	return tailFib(n-1, finalSum, curSum+finalSum)
}
```



### 总结

遇到复杂大问题时，考虑能否分解为相同性质的小问题，若可以再分析分解的终止条件，即可应用递归求解。

分析终止条件时可找最小输入值如 0、1，在其他数据结构中可找最简情形，比如二叉树递归问题的大多数终止条件是叶子节点等等。



### 参考

[廖雪峰的官方网站：递归函数](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/00137473836826348026db722d9435483fa38c137b7e685000)

















