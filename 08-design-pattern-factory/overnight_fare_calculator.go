package designpatternfactory

type OvernightFareCalculator struct {
	fare float64
}

var _ FareCalculator = (*OvernightFareCalculator)(nil)

func NewOvernightFareCalculator() *OvernightFareCalculator {
	return &OvernightFareCalculator{fare: 3.9}
}

func (n *OvernightFareCalculator) Calculate(segment Segment) float64 {
	return segment.distance * n.fare
}
