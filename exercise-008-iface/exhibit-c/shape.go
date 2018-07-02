package main

import (
	"fmt"
	"math"
)

////////////
// Square //
////////////

type Square struct {
	side float64
}

func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

//////////////
// Triangle //
//////////////

type Triangle struct {
	base float64
	height float64
}

func (t *Triangle) Name() string {
	return "Triangle"
}

func (t *Triangle) Perimeter() float64 {
	a := math.Sqrt(math.Pow(t.base/2.0, 2.0) + math.Pow(t.height, 2.0))
	return t.base + 2.0 * a
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

/////////////
// Line    //
/////////////

type Line struct {
	length float64
}

func (l *Line) Name() string {
	return "Line"
}

func (l *Line) Perimeter() float64 {
	return l.length
}

func (l *Line) Area() float64 {
	return 0
}

/////////////
// Hexagon //
/////////////

type Hexagon struct {
	side float64
}

func (h *Hexagon) Name() string {
	return "Hexagon"
}

func (h *Hexagon) Perimeter() float64 {
	return 6 * h.side
}

func (h *Hexagon) Area() float64 {
	return 3 * math.Sqrt(3) * h.side * h.side / 2
}


/////////////
// Decagon //
/////////////

type Decagon struct {
	side float64
}

func (d *Decagon) Name() string {
	return "Decagon"
}

func (d *Decagon) Perimeter() float64 {
	return 10 * d.side
}

func (d *Decagon) Area() float64 {
	return (5.0 * math.Pow(d.side, 2) * math.Sqrt(5.0 + 2 * math.Sqrt(5.0))) / 2.0
}

////////////
// Circle //
////////////

type Circle struct {
	radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

////////////////
// Efficiency //
////////////////

type Shape interface {
	Name() string
	Perimeter() float64
	Area() float64
}

func Efficiency(s Shape) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {
	l := Line{length: 10.0}
	Efficiency(&l)

	t := Triangle{base: 10.0, height: 5.0}
	Efficiency(&t)

	s := Square{side: 10.0}
	Efficiency(&s)

	h := Hexagon{side: 10.0}
	Efficiency(&h)

	d := Decagon{side: 10.0}
	Efficiency(&d)

	c := Circle{radius: 10.0}
	Efficiency(&c)
}
