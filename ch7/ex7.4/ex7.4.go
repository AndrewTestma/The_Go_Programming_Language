package ex7_4

import "io"

//练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）

type stringReader struct {
	s string
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, s.s)
	s.s = s.s[n:]
	if len(s.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
