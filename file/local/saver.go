package local

import (
	"github.com/simp7/nonogram/file"
)

type saver struct {
	path      customPath
	formatter file.Formatter
}

func (s *saver) Save(i interface{}) error {

	err := s.formatter.Encode(i)
	if err != nil {
		return err
	}

	return writeFile(s.path, s.formatter.Content())

}
