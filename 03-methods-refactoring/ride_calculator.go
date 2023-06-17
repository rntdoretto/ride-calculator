package methodsrefactoring

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
		if isValidDistance(segment.Distance) {
			if isValidDateTime(segment.DateTime) {
				if isOvernight(segment.DateTime) {
					if !isSunday(segment.DateTime) {
						fare = fare + segment.Distance*OVERNIGHT_FARE
					} else {
						fare = fare + segment.Distance*OVERNIGHT_SUNDAY_FARE
					}
				} else {
					if isSunday(segment.DateTime) {
						fare = fare + segment.Distance*SUNDAY_FARE
					} else {
						fare = fare + segment.Distance*NORMAL_FARE
					}
				}
			} else {
				return -2
			}
		} else {
			return -1
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
