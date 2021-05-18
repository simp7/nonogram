package standard

type problemUnit struct {
	data [][]int
	max  int
}

func newUnit(data [][]int, max int) problemUnit {
	return problemUnit{data, max}
}

func (u problemUnit) Get(idx int) []int {
	return u.data[idx]
}

func (u problemUnit) Max() int {
	return u.max
}
