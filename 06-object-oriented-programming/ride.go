package objectorientedprogramming

import "time"

const (
	MIN_FARE              = 10
	NORMAL_FARE           = 2.1
	SUNDAY_FARE           = 2.9
	OVERNIGHT_FARE        = 3.9
	OVERNIGHT_SUNDAY_FARE = 5
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
			fare = fare + segment.distance*OVERNIGHT_FARE
		}
		if segment.isOvernight() && segment.isSunday() {
			fare = fare + segment.distance*OVERNIGHT_SUNDAY_FARE
		}
		if !segment.isOvernight() && segment.isSunday() {
			fare = fare + segment.distance*SUNDAY_FARE
		}
		if !segment.isOvernight() && !segment.isSunday() {
			fare = fare + segment.distance*NORMAL_FARE
		}

	}
	if fare < MIN_FARE {
		return MIN_FARE, nil
	} else {
		return fare, nil
	}
}
