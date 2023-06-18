package designpatternfactory

type SundayFareCalculator struct {
	fare float64
}

var _ FareCalculator = (*SundayFareCalculator)(nil)

func NewSundayFareCalculator() *SundayFareCalculator {
	return &SundayFareCalculator{fare: 2.9}
}

func (n *SundayFareCalculator) Calculate(segment Segment) float64 {
	return segment.distance * n.fare
}
