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
	data unit.Map
	raw  []byte
}

func newFormatter() *formatter {
	f := new(formatter)
	f.data = Prototype()
	f.raw = make([]byte, 0)
	return f
}

func (f *formatter) Encode(i interface{}) error {

	switch i.(type) {
	case *nonomap:
		f.data = i.(*nonomap)
		f.raw = convert(i.(*nonomap))
		return nil
	default:
		return invalidFormat
	}

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

	switch i.(type) {
	case unit.Map:
		f.data = i.(unit.Map)
		return nil
	default:
		return invalidFormat
	}

}

func (f *formatter) Raw(content []byte) error {

	f.raw = content

	data := string(content)

	data = strings.TrimSpace(data)
	elements := strings.Split(data, "/")

	width, err := f.getHeight(elements[0])
	if err != nil {
		return err
	}

	height, err := f.getWidth(elements[1])
	if err != nil {
		return err
	}

	bitmap, err := convertToBitmap(width, height, elements[2:])
	if err != nil {
		return err
	}

	f.data = f.data.Init(bitmap)
	return f.data.CheckValidity()

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
