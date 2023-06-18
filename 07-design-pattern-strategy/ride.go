package designpatternstrategy

import (
	"time"
)

const (
	MIN_FARE = 10
)

type Ride struct {
	segments []Segment
}

func NewRide() *Ride {
	return &Ride{
		segments: []Segment{},
	}
}

func (r *Ride) AddSegment(distance float64, dateTime time.Time) error {
	segment, err := NewSegment(distance, dateTime)
	if err != nil {
		return err
	}
	r.segments = append(r.segments, *segment)
	return nil
}

func (r *Ride) CalculateFare() (float64, error) {
	var fare float64
	for _, segment := range r.segments {
		if segment.isOvernight() && !segment.isSunday() {
			fareCalculator := NewOvernightFareCalculator()
			fare = fare + fareCalculator.Calculate(segment)
		}
		if segment.isOvernight() && segment.isSunday() {
			fareCalculator := NewOvernightSundayFareCalculator()
			fare = fare + fareCalculator.Calculate(segment)
		}
		if !segment.isOvernight() && segment.isSunday() {
			fareCalculator := NewSundayFareCalculator()
			fare = fare + fareCalculator.Calculate(segment)
		}
		if !segment.isOvernight() && !segment.isSunday() {
			fareCalculator := NewNormalFareCalculator()
			fare = fare + fareCalculator.Calculate(segment)
		}
	}
	if fare < MIN_FARE {
		return MIN_FARE, nil
	} else {
		return fare, nil
	}
}
