package figures

type Square struct {
	Side float32
}

func (s *Square) getArea() float32 {
	return s.Side * s.Side
}

func (s *Square) getPerimeter() float32 {
	return 4 * s.Side
}
