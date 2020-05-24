package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perim() float64
}

//Defining rectangle and exporting the value Rec
type Rect struct {
	width, height float64
}

//Type square exported
type Squa struct {
	side
}

//Circle type exported
type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}

func (s Squa) Area() float64 {
	return side * side
}

func (s Squa) Perim() float64 {
	return 4 * side
}

func Measure(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perim())
}
