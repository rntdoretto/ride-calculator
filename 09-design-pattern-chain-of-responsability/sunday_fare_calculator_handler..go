package designpatternchainofresponsability

import "fmt"

type SundayFareCalculatorHandler struct {
	fareCalculatorHandler FareCalculatorHandler
	fare                  float64
}

var _ FareCalculatorHandler = (*SundayFareCalculatorHandler)(nil)

func NewSundayFareCalculatorHandler() *SundayFareCalculatorHandler {
	return &SundayFareCalculatorHandler{
		fareCalculatorHandler: nil,
		fare:                  2.9,
	}
}

func NewSundayFareCalculatorHandlerWithNext(fareCalculatorHandler FareCalculatorHandler) *SundayFareCalculatorHandler {
	return &SundayFareCalculatorHandler{
		fareCalculatorHandler: fareCalculatorHandler,
		fare:                  2.9,
	}
}

func (n *SundayFareCalculatorHandler) Calculate(segment Segment) (float64, error) {
	if !segment.isOvernight() && segment.isSunday() {
		return segment.distance * n.fare, nil
	}
	if n.next() == nil {
		return 0, fmt.Errorf("end of chain")
	}
	return n.next().Calculate(segment)
}

func (n *SundayFareCalculatorHandler) next() FareCalculatorHandler {
	return n.fareCalculatorHandler
}
