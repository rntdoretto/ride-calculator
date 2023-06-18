package designpatternchainofresponsability

import "fmt"

type OvernightSundayFareCalculatorHandler struct {
	fareCalculatorHandler FareCalculatorHandler
	fare                  float64
}

var _ FareCalculatorHandler = (*OvernightSundayFareCalculatorHandler)(nil)

func NewOvernightSundayFareCalculatorHandler() *OvernightSundayFareCalculatorHandler {
	return &OvernightSundayFareCalculatorHandler{
		fareCalculatorHandler: nil,
		fare:                  5.0,
	}
}

func NewOvernightSundayFareCalculatorHandlerWithNext(fareCalculatorHandler FareCalculatorHandler) *OvernightSundayFareCalculatorHandler {
	return &OvernightSundayFareCalculatorHandler{
		fareCalculatorHandler: fareCalculatorHandler,
		fare:                  5.0,
	}
}

func (n *OvernightSundayFareCalculatorHandler) Calculate(segment Segment) (float64, error) {
	if segment.isOvernight() && segment.isSunday() {
		return segment.distance * n.fare, nil
	}
	if n.next() == nil {
		return 0, fmt.Errorf("end of chain")
	}
	return n.next().Calculate(segment)
}

func (n *OvernightSundayFareCalculatorHandler) next() FareCalculatorHandler {
	return n.fareCalculatorHandler
}
