package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

// End() - Gets the end point of a square having the starting point
func (s Square) End() Point {
	s.start.x += int(s.a)
	s.start.y += int(s.a)
	newPoint := Point{s.start.x, s.start.y}
	return newPoint
}

// Area() - Gets the area of a Square
func (s Square) Area() uint {
	return s.a * s.a
}

// Perimeter() - Gets the perimeter of a Square
func (s Square) Perimeter() uint {
	return s.a * 4
}
