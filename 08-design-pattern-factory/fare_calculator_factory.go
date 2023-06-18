package designpatternfactory

func CreateFareCalculator(segment Segment) FareCalculator {
	if segment.isOvernight() && !segment.isSunday() {
		return NewOvernightFareCalculator()
	}
	if segment.isOvernight() && segment.isSunday() {
		return NewOvernightSundayFareCalculator()
	}
	if !segment.isOvernight() && segment.isSunday() {
		return NewSundayFareCalculator()
	}
	if !segment.isOvernight() && !segment.isSunday() {
		return NewNormalFareCalculator()
	}
	return nil
}
