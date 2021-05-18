package standard

import (
	"github.com/simp7/nonogram/unit"
)

type problem struct {
	horizontal problemUnit
	vertical   problemUnit
}

func newProblem(hProblem [][]int, vProblem [][]int, hMax int, vMax int) problem {
	return problem{newUnit(hProblem, hMax), newUnit(vProblem, vMax)}
}

func (p problem) Horizontal() unit.ProblemUnit {
	return p.horizontal
}

func (p problem) Vertical() unit.ProblemUnit {
	return p.vertical
}
