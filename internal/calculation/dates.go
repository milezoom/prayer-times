package calculation

import "time"

//GetJulianDay return julian day value for given time struct
func GetJulianDay(t time.Time) float64 {
	year := float64(t.Year())
	month := float64(t.Month())
	date := float64(t.Day())

	if month <= 2 {
		year = year - 1
		month = month + 12
	}

	bVal := 2 + ((year / 100) / 4) - (year / 100)

	return 1720994.5 + (365.25 * year) + (30.6001 * (month + 1)) + bVal + date
}
