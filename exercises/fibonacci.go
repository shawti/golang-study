/*
 * 练习：斐波纳契闭包
 * 现在来通过函数做些有趣的事情。
 *
 * 实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
 * 注：连续的斐波纳契数（Fib[n] = Fib[n-1] + Fib[n-2]）
 */

package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		ret := sum1 + sum2
		sum1 = sum2
		sum2 = ret
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
