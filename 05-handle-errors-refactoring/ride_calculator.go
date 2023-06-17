package handleerrorsrefactoring

import (
	"fmt"
	"time"
)

const (
	MIN_FARE              = 10
	NORMAL_FARE           = 2.1
	SUNDAY_FARE           = 2.9
	OVERNIGHT_FARE        = 3.9
	OVERNIGHT_SUNDAY_FARE = 5
)

type Segment struct {
	Distance float64
	DateTime time.Time
}

func CalculateRide(segments []Segment) (float64, error) {
	var fare float64
	for _, segment := range segments {
		if !isValidDistance(segment.Distance) {
			return 0, fmt.Errorf("invalid distance")
		}
		if !isValidDateTime(segment.DateTime) {
			return 0, fmt.Errorf("invalid date")
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
		return MIN_FARE, nil
	} else {
		return fare, nil
	}
}

func isValidDistance(distance float64) bool {
	return distance > 0
}

func isValidDateTime(dateTime time.Time) bool {
	return !dateTime.IsZero()
}

func isOvernight(dateTime time.Time) bool {
	return dateTime.Hour() >= 22 || dateTime.Hour() <= 6
}

func isSunday(dateTime time.Time) bool {
	return dateTime.Weekday() == time.Sunday
}
