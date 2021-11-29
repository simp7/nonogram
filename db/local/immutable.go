package local

import "github.com/simp7/nonogram/db"

type immutable struct {
	*updater
	*loader
}

func newImmutable(path address, formatter db.Formatter) *immutable {

	i := new(immutable)

	i.updater = &updater{path}
	i.loader = &loader{path, formatter}

	return i

}

func Language(formatter db.Formatter) *immutable {
	return newImmutable(languageAddress(), formatter)
}
