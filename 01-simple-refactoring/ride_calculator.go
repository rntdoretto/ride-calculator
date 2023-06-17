package simplerefactoring

import (
	"reflect"
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

func CalculateRide(segments []map[string]any) float64 {
	var fare float64
	for _, segment := range segments {
		if segment["distance"] != nil && reflect.TypeOf(segment["distance"]).Kind() == reflect.Float64 && segment["distance"].(float64) > 0 {
			if _, ok := segment["dateTime"].(time.Time); segment["dateTime"] != nil && ok && !segment["dateTime"].(time.Time).IsZero() {
				// overnight
				if segment["dateTime"].(time.Time).Hour() >= OVERNIGHT_START || segment["dateTime"].(time.Time).Hour() <= OVERNIGHT_END {
					// not sunday
					if segment["dateTime"].(time.Time).Weekday() != 0 {
						fare = fare + segment["distance"].(float64)*OVERNIGHT_FARE
						// sunday
					} else {
						fare = fare + segment["distance"].(float64)*OVERNIGHT_SUNDAY_FARE
					}
				} else {
					// sunday
					if segment["dateTime"].(time.Time).Weekday() == 0 {
						fare = fare + segment["distance"].(float64)*SUNDAY_FARE
					} else {
						fare = fare + segment["distance"].(float64)*NORMAL_FARE
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
