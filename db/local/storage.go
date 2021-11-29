package local

import (
	"github.com/simp7/nonogram/db"
)

type storage struct {
	*saver
	*loader
}

func newStorage(path address, formatter db.Formatter) *storage {

	s := new(storage)

	s.saver = &saver{path, formatter}
	s.loader = &loader{path, formatter}

	return s

}

func Map(formatter db.Formatter) *storage {
	path := mapAddress()
	return newStorage(path, formatter)
}

func Setting(formatter db.Formatter) *storage {
	path := rootAddress()
	return newStorage(path, formatter)
}
