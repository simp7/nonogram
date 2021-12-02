package standard

import (
	"errors"
	"fmt"
	"github.com/simp7/nonogram/unit"
	"math"
	"strconv"
	"strings"
)

var (
	invalidSize   = errors.New("this file contains invalid size")
	invalidFormat = errors.New("this file doesn't comply with file format")
	invalidMap    = errors.New("map file has been broken")
)

const (
	heightLimit = 30
	widthLimit  = 30
)

type formatter struct {
	data *nonomap
	raw  []byte
}

func newFormatter() *formatter {
	f := new(formatter)
	f.data = Prototype().(*nonomap)
	f.raw = make([]byte, 0)
	return f
}

func (f *formatter) Encode(i interface{}) error {

	if source, ok := i.(*nonomap); ok {
		f.data = source
		f.raw = convert(f.data)
		return nil
	}

	return invalidFormat

}

func convert(nmap *nonomap) []byte {

	result := fmt.Sprintf("%d/%d", nmap.Width(), nmap.Height())

	for _, row := range nmap.bitmap {
		result += fmt.Sprintf("/%d", getRowValue(nmap.Width(), row))
	}

	return []byte(result)

}

func getRowValue(width int, row []bool) int {

	result := 0

	for i, v := range row {
		if v {
			result += int(math.Pow(2, float64(width-i-1)))
		}
	}

	return result

}

func (f *formatter) Decode(i interface{}) error {

	if _, ok := i.(unit.Map); ok {
		origin := i.(*nonomap)
		*origin = *f.data
		return nil
	}

	return invalidFormat

}

func (f *formatter) Raw(content []byte) error {

	f.raw = content

	data := extractData(content)
	bitmap, err := f.getBitmap(data)
	if err != nil {
		return err
	}

	f.data = f.data.Init(bitmap).(*nonomap)
	return f.data.CheckValidity()

}

func extractData(str []byte) []string {
	source := strings.TrimSpace(string(str))
	return strings.Split(source, "/")
}

func (f *formatter) getBitmap(data []string) (result [][]bool, err error) {

	var width, height int

	if width, err = f.getHeight(data[0]); err != nil {
		return
	}

	if height, err = f.getHeight(data[1]); err != nil {
		return
	}

	result, err = convertToBitmap(width, height, data[2:])
	return

}

func convertToBitmap(width, height int, data []string) ([][]bool, error) {

	result := make([][]bool, height)

	for i, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, invalidMap
		}
		result[i] = getBitmapRow(num, width)
	}

	return result, nil

}

func getBitmapRow(value, width int) []bool {

	result := make([]bool, width)

	for i := 1; i <= width; i++ {
		result[width-i] = value%2 == 1
		value = value / 2
	}

	return result

}

func (f *formatter) Content() []byte {
	return f.raw
}

func (f *formatter) Extension() string {
	return "nm"
}

func (f *formatter) getHeight(data string) (height int, err error) {
	height, err = strconv.Atoi(data)
	if !(height > 0 && height <= heightLimit) {
		err = invalidSize
	}
	return
}

func (f *formatter) getWidth(data string) (width int, err error) {
	width, err = strconv.Atoi(data)
	if !(width > 0 && width <= widthLimit) {
		err = invalidSize
	}
	return
}
