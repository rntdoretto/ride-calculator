package original

import (
	"reflect"
	"time"
)

func Calc(movArray []map[string]any) float64 {
	var result float64
	for _, mov := range movArray {
		if mov["dist"] != nil && reflect.TypeOf(mov["dist"]).Kind() == reflect.Float64 && mov["dist"].(float64) > 0 {
			if _, ok := mov["ds"].(time.Time); mov["ds"] != nil && ok && !mov["ds"].(time.Time).IsZero() {

				// overnight
				if mov["ds"].(time.Time).Hour() >= 22 || mov["ds"].(time.Time).Hour() <= 6 {

					// not sunday
					if mov["ds"].(time.Time).Weekday() != 0 {
						result = result + mov["dist"].(float64)*3.90

						// sunday
					} else {
						result = result + mov["dist"].(float64)*5

					}
				} else {
					// sunday
					if mov["ds"].(time.Time).Weekday() == 0 {

						result = result + mov["dist"].(float64)*2.9

					} else {
						result = result + mov["dist"].(float64)*2.10
					}
				}
			} else {
				// fmt.Println(d);
				return -2
			}
		} else {
			// fmt.Println(dist)
			return -1
		}
	}
	if result < 10 {
		return 10
	} else {
		return result
	}
}
