package designpatternchainofresponsability

import "fmt"

type OvernightFareCalculatorHandler struct {
	fareCalculatorHandler FareCalculatorHandler
	fare                  float64
}

var _ FareCalculatorHandler = (*OvernightFareCalculatorHandler)(nil)

func NewOvernightFareCalculatorHandler() *OvernightFareCalculatorHandler {
	return &OvernightFareCalculatorHandler{
		fareCalculatorHandler: nil,
		fare:                  3.9,
	}
}

func NewOvernightFareCalculatorHandlerWithNext(fareCalculatorHandler FareCalculatorHandler) *OvernightFareCalculatorHandler {
	return &OvernightFareCalculatorHandler{
		fareCalculatorHandler: fareCalculatorHandler,
		fare:                  3.9,
	}
}

func (n *OvernightFareCalculatorHandler) Calculate(segment Segment) (float64, error) {
	if segment.isOvernight() && !segment.isSunday() {
		return segment.distance * n.fare, nil
	}
	if n.next() == nil {
		return 0, fmt.Errorf("end of chain")
	}
	return n.next().Calculate(segment)
}

func (n *OvernightFareCalculatorHandler) next() FareCalculatorHandler {
	return n.fareCalculatorHandler
}
