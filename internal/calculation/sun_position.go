package calculation

import (
	"math"
	"time"
)

const earthOrbitCorrection = 0.0167
const earthAxialTilt = -23.44

//GetDayOfYear give day of year of today
func GetDayOfYear() int {
	curDate := time.Now()
	firstDate := time.Date(curDate.Year(), 1, 1, 0, 0, 0, 0, curDate.Location())

	dayNum := curDate.Sub(firstDate).Hours() / 24
	return int(dayNum)
}

//GetAngularDistanceToDate calculate distance in degree for earth to travel in its orbit
func GetAngularDistanceToDate(doy float64) float64 {
	angularSpeedPerDay := 360 / 365.24
	solticeDistance := angularSpeedPerDay * (float64(doy) + 10)
	sinOrbit := math.Sin(DegreeToRadian(angularSpeedPerDay * (float64(doy) - 2)))
	result := solticeDistance + ((360 / math.Pi) * earthOrbitCorrection * sinOrbit)
	return result
}

//GetSunDeclination return declination of sun in degree for given degree distance to date
func GetSunDeclination(degDist float64) float64 {
	return RadianToDegree(
		math.Asin(
			math.Sin(DegreeToRadian(earthAxialTilt)) *
				math.Cos(DegreeToRadian(degDist))))
}

//GetEquationOfTime return equation of time of day in minutes for given dayOfYear and angular distance to date
func GetEquationOfTime(dayOfYear float64, angDist float64) float64 {
	angularSpeed := 360 / 365.24
	solticeDist := angularSpeed * (dayOfYear + 10)
	tanVal := math.Tan(DegreeToRadian(angDist))
	cosVal := math.Cos(DegreeToRadian(-earthAxialTilt))
	cVal := (solticeDist - RadianToDegree(math.Atan(tanVal/cosVal))) / 180
	return 720 * (cVal - math.Round(cVal))
}
