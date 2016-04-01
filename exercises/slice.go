/*
 * 练习：slice
 * 实现 `Pic`。它返回一个 slice 的长度 `dy`，和 slice 中每个元素的长度的 8 位无符号整数 `dx`。当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。
 *
 * 图片的实现已经完成。可能用到的函数包括 (x+y)/2 、 x*y 和 `x^y`（使用 math.Pow 计算最后的函数）。
 *
 * （需要使用循环来分配 [][]uint8 中的每个 `[]uint8`。）

 * （使用 uint8(intValue) 在类型之间进行转换。）
 */

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := range ret {
		ret[i] = make([]uint8, dx)
		for j := range ret[i] {
			ret[i][j] = uint8(i * j)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
