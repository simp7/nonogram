package formatter

import "bytes"

type Buffer struct {
	bytes.Buffer
}

func Raw() *Buffer {
	return &Buffer{bytes.Buffer{}}
}

func (b *Buffer) Raw(from []byte) error {
	_, err := b.Buffer.Write(from)
	return err
}

func (b *Buffer) Content() []byte {
	return b.Buffer.Bytes()
}

func (b *Buffer) Write(data []byte) (int, error) {
	return b.Buffer.Write(data)
}

func (b *Buffer) Read(data []byte) (int, error) {
	return b.Buffer.Read(data)
}
