package nonogram

import "github.com/simp7/nonogram/db"

type System interface {
	Map() db.Storage
	Setting() db.Storage
	Language() db.Immutable
	Maps() ([]string, error)
}
