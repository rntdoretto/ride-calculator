package designpatternchainofresponsability

import "fmt"

type NormalFareCalculatorHandler struct {
	fareCalculatorHandler FareCalculatorHandler
	fare                  float64
}

var _ FareCalculatorHandler = (*NormalFareCalculatorHandler)(nil)

func NewNormalFareCalculatorHandler() *NormalFareCalculatorHandler {
	return &NormalFareCalculatorHandler{
		fareCalculatorHandler: nil,
		fare:                  2.1,
	}
}

func NewNormalFareCalculatorHandlerWithNext(fareCalculatorHandler FareCalculatorHandler) *NormalFareCalculatorHandler {
	return &NormalFareCalculatorHandler{
		fareCalculatorHandler: fareCalculatorHandler,
		fare:                  2.1,
	}
}

func (n *NormalFareCalculatorHandler) Calculate(segment Segment) (float64, error) {
	if !segment.isOvernight() && !segment.isSunday() {
		return segment.distance * n.fare, nil
	}
	if n.next() == nil {
		return 0, fmt.Errorf("end of chain")
	}
	return n.next().Calculate(segment)
}

func (n *NormalFareCalculatorHandler) next() FareCalculatorHandler {
	return n.fareCalculatorHandler
}
