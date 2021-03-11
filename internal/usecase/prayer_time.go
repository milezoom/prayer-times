package usecase

import (
	"fmt"

	"github.com/milezoom/prayer-times/internal/calculation"
)

const lon = 106.903337
const lat = -6.2845348
const alt = 15
const tz = 7

//CalculateDhuhr return hour of dhuhr in date type float64
func CalculateDhuhr(eqTime float64) string {
	dayOfYear := calculation.GetDayOfYear()                             //today date of year
	degDist := calculation.GetAngularDistanceToDate(float64(dayOfYear)) //angular distance to today in degree
	eot := calculation.GetEquationOfTime(float64(dayOfYear), degDist)   //today equation of time

	dhuhrHour := (12 + tz - ((lon / 15) + (eot / 60))) * 3600
	hour := int(dhuhrHour) / 3600
	minute := (int(dhuhrHour) % 3600) / 60
	second := (int(dhuhrHour) % 3600) % 60
	return fmt.Sprint(hour, ":", minute, ":", second)
}
