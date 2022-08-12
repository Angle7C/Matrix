package inter

import (
	"math"
)

const UnitAngle float64 = math.Pi / 180

type I_F interface {
	int64 | float64
}
