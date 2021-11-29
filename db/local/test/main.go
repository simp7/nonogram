package main

import (
	"fmt"
	"github.com/simp7/nonogram/db/local"
	"github.com/simp7/nonogram/unit/standard"
)

func main() {
	s2 := local.Map(standard.Prototype().GetFormatter())

	l, err := s2.List()
	fmt.Println(l)
	fmt.Println(err)
}
