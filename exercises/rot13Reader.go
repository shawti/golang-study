/*
 * 练习：rot13Reader
 * 一个常见模式是 io.Reader 包裹另一个 `io.Reader`，然后通过某种形式修改数据流。
 *
 * 例如，gzip.NewReader 函数接受 `io.Reader`（压缩的数据流）并且返回同样实现了 io.Reader 的 `*gzip.Reader`（解压缩后的数据流）。
 *
 * 编写一个实现了 io.Reader 的 `rot13Reader`， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。
 *
 * 已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 `io.Reader`。
 */

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) rot13(b byte, basis byte) byte {
	return (b-basis+13)%26 + basis
}

func (rot rot13Reader) cryptByte(b byte) byte {

	if b >= 'A' && b <= 'Z' {
		return rot.rot13(b, 'A')
	}

	if b >= 'a' && b <= 'z' {
		return rot.rot13(b, 'a')
	}

	return b
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	l, err := rot.r.Read(b)

	for i := 0; i < l; i++ {
		b[i] = rot.cryptByte(b[i])
	}

	return l, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
