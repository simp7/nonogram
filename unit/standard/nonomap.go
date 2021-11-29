package standard

import (
	"github.com/simp7/nonogram/unit"
)

type nonomap struct {
	width  int
	height int
	bitmap [][]bool
}

/*
	nonomap is divided into 3 parts and has arguments equal or more than 3, which is separated by '/'.

	First two elements indicates width and height respectively.

	Rest elements indicates actual map which player has to solve.
	Each elements indicates map data of each line.
	They are designated by bitmap, which 0 is blank and 1 is filled one.

	Since the size of int is 32bits, width of maps can be equal or less than 32 mathematically.
	But because of display's limit, width and height can't be more than 25

	When it comes to player's map, 2 is checked one where player thinks that cell is blank.

	The extension of file is nm(*.nm)
*/

//Prototype returns prototype of nonogram.Map in this package.
func Prototype() unit.Map {
	return new(nonomap)
}

func (nm *nonomap) Init(bitmap [][]bool) unit.Map {

	result := new(nonomap)

	result.height = len(bitmap)
	result.width = len(bitmap[0])
	result.bitmap = bitmap

	return result

}

func (nm *nonomap) ShouldFilled(x int, y int) bool {
	return nm.bitmap[y][x]
}

func getMaxLength(data [][]int) int {
	max := 0
	for _, v := range data {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func (nm *nonomap) createHorizontalProblemData() [][]int {

	horizontal := make([][]int, nm.height)

	for i := 0; i < nm.height; i++ {

		previousCell := false
		tmp := 0

		for j := 0; j < nm.width; j++ {

			if nm.bitmap[i][j] {
				tmp++
				previousCell = true
			} else {
				if previousCell {
					horizontal[i] = append(horizontal[i], tmp)
					tmp = 0
				}
				previousCell = false
			}

		}

		if previousCell {
			horizontal[i] = append(horizontal[i], tmp)
		}

		if len(horizontal[i]) == 0 {
			horizontal[i] = append(horizontal[i], 0)
		}

	}

	return horizontal

}

func (nm *nonomap) createVerticalProblemData() [][]int {

	vertical := make([][]int, nm.width)

	for i := 0; i < nm.width; i++ {

		previousCell := false
		tmp := 0

		for j := 0; j < nm.height; j++ {
			if nm.bitmap[j][i] {
				tmp++
				previousCell = true
			} else {
				if previousCell {
					vertical[i] = append(vertical[i], tmp)
					tmp = 0
				}
				previousCell = false
			}
		}

		if previousCell {
			vertical[i] = append(vertical[i], tmp)
		}

		if len(vertical[i]) == 0 {
			vertical[i] = append(vertical[i], 0)
		}

	}

	return vertical

}

func (nm *nonomap) CreateProblem() unit.Problem {

	hData := nm.createHorizontalProblemData()
	vData := nm.createVerticalProblemData()

	hMax := getMaxLength(hData)
	vMax := getMaxLength(vData)

	return newProblem(hData, vData, hMax, vMax)

}

//This function returns height of nonomap

func (nm *nonomap) Height() int {
	return nm.height
}

func (nm *nonomap) Width() int {
	return nm.width
}

func (nm *nonomap) FilledTotal() (total int) {

	total = 0

	for n := range nm.bitmap {
		total += nm.countRow(n)
	}

	return

}

func (nm *nonomap) countRow(y int) int {
	result := 0
	for _, v := range nm.bitmap[y] {
		if v {
			result++
		}
	}
	return result
}

func (nm *nonomap) HeightLimit() int {
	return 30
}

func (nm *nonomap) WidthLimit() int {
	return 30
}

func (nm *nonomap) CheckValidity() error {
	if nm.height > nm.HeightLimit() || nm.width > nm.WidthLimit() || nm.height <= 0 || nm.width <= 0 {
		return invalidMap
	}
	return nil
}

func (nm *nonomap) GetFormatter() unit.Formatter {
	return newFormatter()
}
