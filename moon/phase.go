package moon

import (
	"math"
	"time"
)

const (
	synodicMonth    = 29.53058867
	julianOffset    = 4716.0
	monthCorrection = 30.6001
	julianEpoch     = 1524.5
)

func getPhase(t time.Time) float64 {
	// 6 jan 2000
	const knownNewMoon = 2451550.259

	jd := julianDay(t)

	age := math.Mod(jd-knownNewMoon, synodicMonth)
	if age < 0 {
		age += synodicMonth
	}

	// phase 0.0..1.0
	return age / synodicMonth
}

func julianDay(t time.Time) float64 {
	t = t.UTC()
	y := float64(t.Year())
	m := float64(t.Month())
	d := float64(t.Day()) +
		float64(t.Hour())/24.0 +
		float64(t.Minute())/1440.0 +
		float64(t.Second())/86400.0

	if m <= 2 {
		y--
		m += 12
	}

	a := math.Floor(y / 100)
	b := 2 - a + math.Floor(a/4)

	return math.Floor(365.25*(y+julianOffset)) +
		math.Floor(monthCorrection*(m+1)) +
		d + b - julianEpoch
}
