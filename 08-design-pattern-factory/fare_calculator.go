package designpatternfactory

type FareCalculator interface {
	Calculate(segment Segment) float64
}
