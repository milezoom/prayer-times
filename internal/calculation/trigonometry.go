package calculation

import "math"

//DegreeToRadian convert degree value to radian
func DegreeToRadian(deg float64) float64 {
	return deg * (math.Pi / 180)
}

//RadianToDegree convert radian value to degree
func RadianToDegree(rad float64) float64 {
	return rad * (180 / math.Pi)
}
