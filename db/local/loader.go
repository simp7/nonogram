package local

import (
	"github.com/simp7/nonogram/db"
	"os"
)

type loader struct {
	path      address
	formatter db.Formatter
}

func (l *loader) Load(name string, target interface{}) error {

	data, err := readRealFile(l.path, name+"."+l.formatter.Extension())
	if err != nil {
		return err
	}

	if err = l.formatter.Raw(data); err != nil {
		return err
	}

	return l.formatter.Decode(target)

}

func (l *loader) List() (list []string, err error) {

	realAddr, err := l.path.Real()
	if err != nil {
		return
	}

	files, err := os.ReadDir(realAddr)
	if err != nil {
		return
	}

	for _, v := range files {
		list = append(list, v.Name())
	}

	return

}
