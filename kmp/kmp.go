package kmp

//
// 暴力向后按字符匹配 O(M*N)
//
func ForceTraverse(source, pattern string) []int {
	n, m := len(source), len(pattern)
	var locs []int
	for i := 0; i <= n-m; i++ {
		j := 0
		k := i
		for j < m && pattern[j] == source[k] {
			j++
			k++
			if j == m {
				locs = append(locs, i)
				j = 0
				break
			}
		}
	}
	return locs
}

//
// KMP 解决暴力匹配的重复匹配问题 O(M+N)
// 通过在子串中找到下一个进行匹配的字符，避免回到主串头部进行比较
//
func KMP(source, pattern string) []int {
	var locs []int
	i, j := 0, 0
	next := findNext(pattern) // 生成子串的 next 数组
	n, m := len(source), len(next)

	for i < n {

		if j < m {
			if source[i] == pattern[j] { // 匹配上了
				i++
				j++
			} else if j > 0 {
				j = next[j-1] // 一直匹配直到不匹配，倒回到前缀对应位置处
			} else {
				i++ // 一直没匹配上
			}
		}

		// 发现一个完全匹配
		if j == m {
			locs = append(locs, i-m)
			i++
			j = 0
		}
	}

	return locs
}

//
// 高效完成前缀后缀匹配 O(N)
//
func findNext(s string) []int {
	n := len(s)
	next := make([]int, n)
	j, i := 0, 1

	for i < n {
		if s[i] == s[j] {
			// 一切顺利，继续匹配
			next[i] = j + 1 // j 既是索引，又是计数器
			i++
			j++
		} else if j == 0 {
			// 一直没找到匹配上的
			i++
		} else {
			// 匹配了一会儿中断了，回退一位
			j = next[j-1]
		}
	}
	return next
}
