package cognitivecomplexityrefactoring

import (
	"time"
)

const (
	MIN_FARE              = 10
	NORMAL_FARE           = 2.1
	SUNDAY_FARE           = 2.9
	OVERNIGHT_FARE        = 3.9
	OVERNIGHT_SUNDAY_FARE = 5
	OVERNIGHT_START       = 22
	OVERNIGHT_END         = 6
)

type Segment struct {
	Distance float64
	DateTime time.Time
}

func CalculateRide(segments []Segment) float64 {
	var fare float64
	for _, segment := range segments {
		if !isValidDistance(segment.Distance) {
			return -1
		}
		if !isValidDateTime(segment.DateTime) {
			return -2
		}
		if isOvernight(segment.DateTime) && !isSunday(segment.DateTime) {
			fare = fare + segment.Distance*OVERNIGHT_FARE
		}
		if isOvernight(segment.DateTime) && isSunday(segment.DateTime) {
			fare = fare + segment.Distance*OVERNIGHT_SUNDAY_FARE
		}
		if !isOvernight(segment.DateTime) && isSunday(segment.DateTime) {
			fare = fare + segment.Distance*SUNDAY_FARE
		}
		if !isOvernight(segment.DateTime) && !isSunday(segment.DateTime) {
			fare = fare + segment.Distance*NORMAL_FARE
		}
	}
	if fare < MIN_FARE {
		return MIN_FARE
	} else {
		return fare
	}
}

func isValidDistance(distance float64) bool {
	return distance > 0
}

func isValidDateTime(dateTime time.Time) bool {
	return !dateTime.IsZero()
}

func isOvernight(dateTime time.Time) bool {
	return dateTime.Hour() >= OVERNIGHT_START || dateTime.Hour() <= OVERNIGHT_END
}

func isSunday(dateTime time.Time) bool {
	return dateTime.Weekday() == time.Sunday
}
