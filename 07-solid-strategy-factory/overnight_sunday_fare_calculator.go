package solidstrategyfactory

type OvernightSundayFareCalculator struct {
	fare float64
}

var _ FareCalculator = (*OvernightSundayFareCalculator)(nil)

func NewOvernightSundayFareCalculator() *OvernightSundayFareCalculator {
	return &OvernightSundayFareCalculator{fare: 5.0}
}

func (n *OvernightSundayFareCalculator) calculate(segment Segment) float64 {
	return segment.distance * n.fare
}
