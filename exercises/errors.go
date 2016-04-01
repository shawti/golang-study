/*
 *练习：错误
 * 从之前的练习中复制 Sqrt 函数，并修改使其返回 error 值。
 *
 * Sqrt 接收到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。
 *
 * 创建一个新类型

 * type ErrNegativeSqrt float64
 * 为其实现

 * func (e ErrNegativeSqrt) Error() string
 * 使其成为一个 `error`， 该方法就可以让 ErrNegativeSqrt(-2).Error() 返回 `"cannot Sqrt negative number: -2"`。

 * 注意： 在 Error 方法内调用 fmt.Sprint(e) 将会让程序陷入死循环。可以通过先转换 e 来避免这个问题：`fmt.Sprint(float64(e))`。请思考这是为什么呢？

 * 修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
 */

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	const E = 1e-10
	z := float64(1)
	k := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if (z-k <= E && z-k >= -E) || z == k {
			return z, nil
		}
		k = z
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
