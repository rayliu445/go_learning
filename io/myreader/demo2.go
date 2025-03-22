package myreader

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
