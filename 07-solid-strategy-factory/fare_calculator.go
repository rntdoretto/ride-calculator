package solidstrategyfactory

type FareCalculator interface {
	calculate(segment Segment) float64
}
