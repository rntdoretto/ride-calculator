package designpatternstrategy

type NormalFareCalculator struct {
	fare float64
}

var _ FareCalculator = (*NormalFareCalculator)(nil)

func NewNormalFareCalculator() *NormalFareCalculator {
	return &NormalFareCalculator{fare: 2.1}
}

func (n *NormalFareCalculator) Calculate(segment Segment) float64 {
	return segment.distance * n.fare
}
