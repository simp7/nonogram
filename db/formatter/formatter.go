package formatter

import (
	"github.com/simp7/nonogram/db"
	"strings"
)

func GetFormat(fileName string) db.Formatter {
	elements := strings.Split(fileName, ".")
	switch elements[1] {
	case "json":
		return Json()
	case "yaml":
		return Yaml()
	default:
		return nil
	}
}
