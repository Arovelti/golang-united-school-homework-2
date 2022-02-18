package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type customInt int

var SidesTriangle customInt = 3
var SidesSquare customInt = 4
var SidesCircle customInt = 0

func CalcSquare(sideLen float64, sidesNum customInt) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3 * sideLen * sideLen)
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0.0
	}

}
