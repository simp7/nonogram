package local

import (
	"fmt"
	"github.com/simp7/nonogram/file"
	"os"
	"strings"
)

type mapList struct {
	unit        int
	dirPath     []byte
	files       []os.DirEntry
	currentFile string
	currentPage int
}

func newMapList(by int) file.MapList {

	list := new(mapList)

	list.unit = by
	list.currentPage = 1
	list.Refresh()

	return list

}

/*
	This function returns list of map whose number of maps are separated by 10.
	This function will be called when player enter the select page.
*/

func (l *mapList) Current() []string {

	list := make([]string, l.unit)

	for n := 0; n < l.unit; n++ {
		order := l.realIdx(n)
		if order < len(l.files) {
			fileName := l.files[order].Name()
			list[n] = fmt.Sprintf("%d. %s", n, trimSuffix(fileName))
		}
	}

	return list

}

/*
	This function gets player to the next page of list.
	This function will be called when player inputs left-arrow key.
*/

func (l *mapList) Next() {
	if l.unit*l.currentPage >= len(l.files) {
		l.currentPage = 1
	} else {
		l.currentPage++
	}
}

/*
	This function gets player to the previous page of list
	This function will be called when player inputs right-arrow key.
*/

func (l *mapList) Prev() {
	if l.currentPage == 1 {
		l.currentPage = (len(l.files) - 1) / l.unit
	} else {
		l.currentPage--
	}
}

/*
	This function returns player's current page.
	This function will be called with list of map, attached with list header.
*/

func (l *mapList) CurrentPage() int {
	return l.currentPage
}

func (l *mapList) LastPage() int {
	return len(l.files)/l.unit + 1
}

/*
	This function gets nonomap data by number.
	This function will be called when user inputs number in select.
*/

func (l *mapList) GetMapName(idx int) (string, bool) {

	if idx >= len(l.files) {
		return "", false
	}

	l.currentFile = trimSuffix(l.files[l.realIdx(idx)].Name())
	return l.currentFile, true

}

func (l *mapList) GetCachedMapName() string {
	return l.currentFile
}

/*
	This function refresh list of map
	This function will be called after user create map so it contains added map.
*/

func (l *mapList) Refresh() error {

	mapDir, err := get(MAPSDIR)
	if err != nil {
		return err
	}

	l.files, err = readDir(mapDir)
	return err

}

func (l *mapList) realIdx(idx int) int {
	return idx + l.unit*(l.currentPage-1)
}

func trimSuffix(name string) string {
	return strings.TrimSuffix(name, ".nm")
}
