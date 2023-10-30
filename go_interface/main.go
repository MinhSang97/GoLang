package main

import (
	"app/app/myinterfaces"
	"fmt"
)

type Triangle struct {
	Height int64
	Widght int64
}

func (t Triangle) GetArea() float64 {
	return float64(t.Height*t.Widght) * 0.5
}
func main() {
	a := Triangle{2, 3}
	fmt.Println(a.GetArea())
	PrintArea(a)
}

func PrintArea(s myinterfaces.Shape) {
	fmt.Println(s.GetArea())
}
