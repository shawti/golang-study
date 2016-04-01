/*
 * 练习：map
 * 实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。
 *
 * 你会发现 strings.Fields 很有帮助。
 */

package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	strArr := strings.Fields(s)
	for _, v := range strArr {
		ret[v]++
	}

	return ret
}

func main() {
	wc.Test(WordCount)
}
