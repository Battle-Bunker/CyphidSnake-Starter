package lib

import (
	"github.com/BattlesnakeOfficial/rules"
	"math"
)

func ManhattanDistance(p1, p2 rules.Point) float64 {
	return math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
}
