/*
 * 练习：循环和函数
 * 作为练习函数和循环的简单途径，用牛顿法实现开方函数。
 *
 * 在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值：
 *
 * z = z - (z*z-x)/(2*z)
 *
 * 为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数是否变化。结果与 [[http://golang.org/pkg/math/#Sqrt][math.Sqrt] 接近吗？
 *
 * 提示：定义并初始化一个浮点值，向其提供一个浮点语法或使用转换：
 *
 * z := float64(1)
 * z := 1.0
 */

package main

import (
	"fmt"
	"math"
)

// Sqrt returns its argument float64 by sqrt.
func Sqrt(x float64) float64 {
	const E = 1e-10
	z := float64(1)
	k := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if (z-k <= E && z-k >= -E) || z == k {
			return z
		}
		k = z
	}
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
}
