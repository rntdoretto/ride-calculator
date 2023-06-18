package designpatternstrategy

import (
	"fmt"
	"time"
)

const (
	OVERNIGHT_START = 22
	OVERNIGHT_END   = 6
)

type Segment struct {
	distance float64
	dateTime time.Time
}

func NewSegment(distance float64, dateTime time.Time) (*Segment, error) {
	if !isValidDistance(distance) {
		return nil, fmt.Errorf("invalid distance")
	}
	if !isValidDateTime(dateTime) {
		return nil, fmt.Errorf("invalid date")
	}
	return &Segment{distance, dateTime}, nil
}

func isValidDistance(distance float64) bool {
	return distance > 0
}

func isValidDateTime(dateTime time.Time) bool {
	return !dateTime.IsZero()
}

func (s *Segment) isOvernight() bool {
	return s.dateTime.Hour() >= OVERNIGHT_START || s.dateTime.Hour() <= OVERNIGHT_END
}

func (s *Segment) isSunday() bool {
	return s.dateTime.Weekday() == time.Sunday
}
