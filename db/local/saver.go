package local

import (
	"github.com/simp7/nonogram/db"
)

type saver struct {
	path      address
	formatter db.Formatter
}

func (s *saver) Save(name string, i interface{}) error {

	if err := s.formatter.Encode(i); err != nil {
		return err
	}

	return writeFile(s.path, name+"."+s.formatter.Extension(), s.formatter.Content())

}
