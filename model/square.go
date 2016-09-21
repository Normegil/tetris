package model

import "fmt"

type Square struct {
	X, Y  int
	Color Color
}

func (s Square) String() string {
	return fmt.Sprintf("{X:%d, Y:%d}", s.X, s.Y)
}
