package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	X      int
	Y      int
	Length int
}

// NewSquare returns a new square
func NewSquare(x, y, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length cannot be zero or negative")
	}
	s := Square{ /* values could be specified here */ }
	s.X = x
	s.Y = y
	s.Length = length
	return &s, nil
}

// Move moves the square
func (s *Square) Move(dx, dy int) {
	s.X += dx
	s.Y += dy
}

// Area returns the square area
//      ⬐ pointer is not necessary here, Square is not changed
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatal("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
