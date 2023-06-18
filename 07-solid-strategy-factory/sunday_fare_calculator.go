package solidstrategyfactory

type SundayFareCalculator struct {
	fare float64
}

var _ FareCalculator = (*SundayFareCalculator)(nil)

func NewSundayFareCalculator() *SundayFareCalculator {
	return &SundayFareCalculator{fare: 2.9}
}

func (n *SundayFareCalculator) calculate(segment Segment) float64 {
	return segment.distance * n.fare
}
