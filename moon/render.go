package moon

import (
	"fmt"
	"math"
	"time"
)

func Render() {
	r := 15
	phase := getPhase(time.Now())
	waxing := phase < 0.5

	var illumination float64
	if waxing {
		illumination = phase / 0.5
	} else {
		illumination = (1.0 - phase) / 0.5
	}

	if illumination < 0.001 {
		illumination = 0.0
	}
	if illumination > 0.999 {
		illumination = 1.0
	}

	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			dist := math.Sqrt(float64(x*x + y*y))
			if dist > float64(r) {
				fmt.Print("  ")
				continue
			}

			halfWidth := math.Sqrt(float64(r*r - y*y))

			tx := halfWidth * (1.0 - 2.0*illumination)

			var lit bool
			if illumination == 0.0 {
				lit = false
			} else if illumination == 1.0 {
				lit = true
			} else if waxing {
				lit = float64(x) >= tx
			} else {
				lit = float64(x) <= tx
			}

			if lit {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

}
