package designpatternchainofresponsability

import (
	"time"
)

const (
	MIN_FARE = 10
)

type Ride struct {
	fareCalculatorHandler FareCalculatorHandler
	segments              []Segment
}

func NewRide(fareCalculatorHandler FareCalculatorHandler) *Ride {
	return &Ride{
		fareCalculatorHandler: fareCalculatorHandler,
		segments:              []Segment{},
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
		f, err := r.fareCalculatorHandler.Calculate(segment)
		if err != nil {
			return MIN_FARE, nil
		}
		fare = fare + f
	}
	if fare < MIN_FARE {
		return MIN_FARE, nil
	} else {
		return fare, nil
	}
}
