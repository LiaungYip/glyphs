package glyphs

import (
	"math"
)

type xy struct {
	X, Y float64
}

type originAndScale struct {
	origin_x, origin_y, scale float64
}

func (fxy xy) moveAndScale(origin_x, origin_y, scale float64) xy {
	x := (fxy.X * scale) + origin_x
	y := (fxy.Y * scale) + origin_y
	return xy{X: x, Y: y}
}

type floatPolar struct {
	Mag, Deg float64
}

// pol2cart converts from a magnitude and angle into cartesian X & Y.
func (fp floatPolar) xy() xy {
	x := math.Cos(fp.Deg/180.0*math.Pi) * fp.Mag
	y := -math.Sin(fp.Deg/180.0*math.Pi) * fp.Mag
	return xy{X: x, Y: y}
}
