package designpatternchainofresponsability

type FareCalculatorHandler interface {
	Calculate(segment Segment) (float64, error)
	next() FareCalculatorHandler
}
