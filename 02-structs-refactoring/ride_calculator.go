package structsrefactoring

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
		if segment.Distance > 0 {
			if !segment.DateTime.IsZero() {
				// overnight
				if segment.DateTime.Hour() >= OVERNIGHT_START || segment.DateTime.Hour() <= OVERNIGHT_END {
					// not sunday
					if segment.DateTime.Weekday() != time.Sunday {
						fare = fare + segment.Distance*OVERNIGHT_FARE
						// sunday
					} else {
						fare = fare + segment.Distance*OVERNIGHT_SUNDAY_FARE
					}
				} else {
					// sunday
					if segment.DateTime.Weekday() == time.Sunday {
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
