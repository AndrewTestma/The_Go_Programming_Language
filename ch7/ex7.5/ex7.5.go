package ex7_5

import (
	"io"
)

//练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。

type limitReader struct {
	r        io.Reader
	n, limit int64
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += int64(n)
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, limit: n}
}
