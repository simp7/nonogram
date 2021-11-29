package formatter

import "bytes"

type Buffer struct {
	buf bytes.Buffer
}

func Raw() Buffer {
	return Buffer{bytes.Buffer{}}
}

func (b Buffer) Raw(from []byte) error {
	_, err := b.buf.Write(from)
	return err
}

func (b Buffer) Content() []byte {
	return b.buf.Bytes()
}

func (b Buffer) Write(data []byte) (int, error) {
	return b.buf.Write(data)
}

func (b Buffer) Read(data []byte) (int, error) {
	return b.buf.Read(data)
}
